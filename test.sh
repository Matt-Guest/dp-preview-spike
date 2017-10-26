URL=http://localhost:9091/preview
URL1=http://localhost:9091/preview1
URL2=http://localhost:9091/preview2
URL3=http://localhost:9091/preview2
URL4=http://localhost:9091/preview2
curl -s -w "%{time_total}\n" "$URL?[1-10000]" -o /dev/null | tee test1.csv &

curl -s -w "%{time_total}\n" "$URL1?[1-10000]" -o /dev/null | tee test2.csv &

curl -s -w "%{time_total}\n" "$URL2?[1-10000]" -o /dev/null | tee test3.csv &

curl -s -w "%{time_total}\n" "$URL3?[1-10000]" -o /dev/null | tee test4.csv &

curl -s -w "%{time_total}\n" "$URL4?[1-10000]" -o /dev/null | tee test5.csv &

