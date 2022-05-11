
docker pull postgres


docker run --name postgres \
--restart always \
-e ALLOW_IP_RANGE=0.0.0.0/0 \
-e POSTGRES_PASSWORD='67c48438-ce78-11ec-a57e-00155dfcb67f' \
-p 5432:5432 \
-v pgdata:/var/lib/postgresql/data \
-d postgres