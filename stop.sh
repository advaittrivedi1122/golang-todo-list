# Stop ScyllaDb in docker
echo -e "\nStopping ScyllaDb Container"
sudo docker container stop scylla-db > /dev/null 2>&1

# Wait for ScyllaDb to be closed properly inside container
echo -e "\nGracefully closing ScyllaDb (wait 15s)"
sleep 15

echo -e "\nRemoving ScyllaDb Container"
sudo docker container rm scylla-db > /dev/null 2>&1
echo -e "\nRemoved ScyllaDb Container Successfully"