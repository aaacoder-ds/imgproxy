#!/bin/bash

case "$IMGPROXY_MALLOC" in

  malloc)
    # Do nothing
    ;;

  jemalloc)
    export LD_PRELOAD="$LD_PRELOAD:/usr/local/lib/libjemalloc.so"
    ;;

  tcmalloc)
    export LD_PRELOAD="$LD_PRELOAD:/usr/local/lib/libtcmalloc_minimal.so"
    ;;

  *)
    echo "Unknown malloc: $IMGPROXY_MALLOC"
    exit 1
esac

# Start cron service in background
echo "Starting cron service..."
service cron start

# Create log file for cleanup script
touch /var/log/cleanup.log
chmod 666 /var/log/cleanup.log

echo "Cron service started. Cleanup will run every 6 hours."
echo "Cleanup logs will be written to /var/log/cleanup.log"

# Execute the original command (imgproxy)
exec "$@" 