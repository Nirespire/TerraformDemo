cd  node-function

node docker-npm.js

cd ..

zip -j function.zip node-function/index.js node-function/package.json node-function/terraform-admin.json
zip -r function.zip ../node_modules