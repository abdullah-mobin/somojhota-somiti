#!/bin/bash
# Start MongoDB in the background
mongod --replSet rs0 --bind_ip_all --port 27019 &

# Wait until MongoDB is ready to accept connections
echo "Waiting for MongoDB to start..."
until mongosh --port 27019 --quiet --eval "db.runCommand({ ping: 1 })" >/dev/null 2>&1; do
  sleep 2
done

echo "MongoDB is up. Checking replica set status..."

# Check if replica set already initialized
IS_RS_INIT=$(mongosh --port 27019 --quiet --eval "rs.status().ok" 2>/dev/null || echo 0)

if [ "$IS_RS_INIT" != "1" ]; then
  echo "Initiating replica set..."
  mongosh --port 27019 --quiet --eval '
    rs.initiate({
      _id: "rs0",
      members: [{ _id: 0, host: "localhost:27019" }]
    })
  '
else
  echo "Replica set already initialized."
fi

# Keep the process running in foreground
wait