export RELEASE=$(curl -s "https://api.github.com/repos/lamarios/unifi-throughput-ncurses/releases/tags/${DRONE_TAG}" |  jq .id)
export FILE="$(ls unifi-throughput-*.tar.gz)"
export FILE_PATH="$(pwd)/$FILE"
echo "Release $RELEASE, File $FILE - $FILE_PATH"
curl --location --request POST "https://uploads.github.com/repos/lamarios/unifi-throughput-ncurses/releases/${RELEASE}/assets?name=${FILE}" --data-binary "@${FILE_PATH}" -H "Content-type: application/tar+gzip" -H "Authorization: token ${GITHUB_TOKEN}"
