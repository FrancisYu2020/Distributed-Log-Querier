cd /home/hangy6/mp1-hangy6-tian23
git add .
git commit -m "new commit"
git push origin hangy6
cd src/scripts/
bash update_source_code.sh

# test 2 or 3 machines on these and expand
bash build_all_server_client.sh
bash kill_server.sh
bash run_all_server.sh