cd ../../
git add .
git commit -m "new commit"
git push
cd src/scripts/
bash update_source_code.sh
bash build_all_server_client.sh
bash kill_server.sh
bash run_all_server.sh