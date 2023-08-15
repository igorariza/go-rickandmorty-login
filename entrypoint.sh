#!/bin/bash
set -e

DB_USER='202mo42z6d93dr3inlxd';
DB_PASSWORD='pscale_pw_b4NXD2us6RVaKa3mudvosw2bY6xJzzz3CTypPF4VzjM';
DB_HOST='aws.connect.psdb.cloud';
DB_PORT=3306;
DB_NAME='apirickandmorty';

TOTAL_PAGES=$(curl -s https://rickandmortyapi.com/api/character | jq -r .info.pages)
i=0
k=0
while [ "$i" -le ${TOTAL_PAGES} ]; do
        i=$((i+1));
                
        curl --request GET -s "https://rickandmortyapi.com/api/character?page=${i}" | jq -r '.results' > res.json
        
        jq -c '.[]' res.json | while read j; do                   
                name=$(echo ${j}  | jq -r '.name')
                status=$(echo ${j}  | jq -r '.status')
                species=$(echo ${j}  | jq -r '.species')
                gender=$(echo ${j}  | jq -r '.gender')
                image=$(echo ${j}  | jq -r '.image')
                created=$(echo ${j}  | jq -r '.created')
                
                echo "INSERT INTO characters (name,status,species,gender,image,created) VALUES ('${name//\'/\'}','${status//\'/\'}','${species}','${gender//\'/\'}','${image//\'/\'}','${created}')"  | mysql --user=$DB_USER --password=$DB_PASSWORD --host=$DB_HOST $DB_NAME 
                
        done    
done

exec "$@"