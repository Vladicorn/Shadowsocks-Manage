# Shadowsocks Manage
 

## Для билда
docker build -t shad .
## Для запуска 
docker run --network host shad

## API
GET http://127.0.0.1:8086/api/create/:port 
### Метод шифрования chacha20-ietf-poly1305