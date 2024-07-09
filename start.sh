# Start ScyllaDb in docker
echo -e "\nStarting ScyllaDb Container"
sudo docker run --name scylla-db -d scylladb/scylla > /dev/null 2>&1

# Wait for ScyllaDb to be initialised properly inside container
echo -e "\nConfiguring ScyllaDb (wait 60s)"
sleep 60


# create keyspace
HOST=172.17.0.2
PORT=9042
KEYSPACE=todolist

echo -e "\nCreating Keyspace for ScyllaDb"

cqlsh $HOST $PORT -e "CREATE KEYSPACE IF NOT EXISTS $KEYSPACE WITH REPLICATION = {'class':'NetworkTopologyStrategy', 'replication_factor':1}" > /dev/null 2>&1

echo -e "\nKeyspace created successfully"

# Create .env from .env.example
cp .env.example .env
echo -e "\nEnvironment variables set successfully"

# Start todoList 
echo -e "\nStarting server"
./todoList