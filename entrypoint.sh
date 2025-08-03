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

# Start cron service in background (as root)
echo "Starting cron service..."
if [ "$(id -u)" = "0" ]; then
  service cron start
  echo "Cron service started. Cleanup will run every 6 hours."
  echo "Cleanup logs will be written to /var/log/cleanup.log"
  
  # Switch to imgproxy user for running the application
  exec gosu imgproxy "$@"
else
  echo "Warning: Cannot start cron service as non-root user. Cleanup will not run automatically."
  echo "You can run cleanup manually with: docker exec -it <container> /usr/local/bin/cleanup-script.sh"
  
  # Execute the original command (imgproxy)
  exec "$@"
fi 