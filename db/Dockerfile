FROM mongo

COPY init.json init.json
CMD mongoimport --username jl --password pwd --host mongodb --db pairot --collection Teams --type json --file init.json --jsonArray