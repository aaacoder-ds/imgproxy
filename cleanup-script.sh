#!/bin/bash

# imgproxy data cleanup script for Docker container
# This script removes old cached images to maintain disk space under 25GB

DATA_DIR="/data"
MAX_SIZE_GB=25

# Check if we're inside a container
if [ ! -f /.dockerenv ]; then
    echo "This script should run inside a Docker container"
    exit 1
fi

# Get current disk usage in GB (using du for container filesystem)
CURRENT_SIZE_GB=$(du -s "$DATA_DIR" 2>/dev/null | awk '{print int($1/1024/1024)}')

if [ -z "$CURRENT_SIZE_GB" ] || [ "$CURRENT_SIZE_GB" -eq 0 ]; then
    echo "$(date): No data directory found or empty. Exiting."
    exit 0
fi

echo "$(date): Starting cleanup process"
echo "Current disk usage: ${CURRENT_SIZE_GB}GB"

# If we're under the limit, no cleanup needed
if [ "$CURRENT_SIZE_GB" -lt "$MAX_SIZE_GB" ]; then
    echo "Disk usage is under ${MAX_SIZE_GB}GB limit. No cleanup needed."
    exit 0
fi

# Calculate how much we need to free up (in GB)
NEED_TO_FREE=$((CURRENT_SIZE_GB - MAX_SIZE_GB + 2))  # Free up extra 2GB for safety
echo "Need to free up approximately ${NEED_TO_FREE}GB"

# Find and remove oldest files until we've freed enough space
FREED_SPACE=0
while [ "$FREED_SPACE" -lt "$NEED_TO_FREE" ]; do
    # Find the oldest file in the data directory (using stat for container compatibility)
    OLDEST_FILE=$(find "$DATA_DIR" -type f -exec stat -c '%Y %n' {} \; | sort -n | head -1 | cut -d' ' -f2-)
    
    if [ -z "$OLDEST_FILE" ] || [ ! -f "$OLDEST_FILE" ]; then
        echo "No more files to remove"
        break
    fi
    
    # Get file size in GB
    FILE_SIZE_BYTES=$(stat -c%s "$OLDEST_FILE" 2>/dev/null)
    if [ -z "$FILE_SIZE_BYTES" ]; then
        echo "Could not get file size for $OLDEST_FILE, removing anyway"
        rm -f "$OLDEST_FILE"
        continue
    fi
    
    FILE_SIZE_GB=$(echo "scale=2; $FILE_SIZE_BYTES / 1024 / 1024 / 1024" | bc -l 2>/dev/null || echo "0")
    
    echo "Removing: $OLDEST_FILE (${FILE_SIZE_GB}GB)"
    rm -f "$OLDEST_FILE"
    
    FREED_SPACE=$(echo "$FREED_SPACE + $FILE_SIZE_GB" | bc -l 2>/dev/null || echo "$FREED_SPACE")
    echo "Total freed: ${FREED_SPACE}GB"
done

# Final disk usage check
FINAL_SIZE_GB=$(du -s "$DATA_DIR" 2>/dev/null | awk '{print int($1/1024/1024)}')
echo "Cleanup completed. Final disk usage: ${FINAL_SIZE_GB}GB"
echo "$(date): Cleanup process finished" 