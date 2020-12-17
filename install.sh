#!/bin/bash

SBCLI_URI="https://github.com/sivaprasadreddy/spring-boot-cli/releases/download"
SBCLI_DIR="$HOME/.sbcli"
SBCLI_VERSION="v0.0.4"

sbcli_archives_folder="${SBCLI_DIR}/archives"
sbcli_zip_file="${sbcli_archives_folder}/spring-boot-cli-${SBCLI_VERSION}.zip"

mkdir -p "$sbcli_archives_folder"
rm -rf "$SBCLI_DIR/$SBCLI_VERSION"
mkdir -p "$SBCLI_DIR/$SBCLI_VERSION"
mkdir -p "$SBCLI_DIR/current"

ostype="linux"
case "$(uname -s)" in
    Darwin*)
        ostype="darwin"
        ;;
    Linux*)
        ostype="linux"
        ;;
esac

echo "Downloading $ostype binary ${SBCLI_URI}/${SBCLI_VERSION}/spring-boot-cli_${ostype}_amd64.zip"
curl --location --progress-bar "${SBCLI_URI}/${SBCLI_VERSION}/spring-boot-cli_${ostype}_amd64.zip" > "$sbcli_zip_file"

tar -xf $sbcli_zip_file -C $sbcli_archives_folder "spring-boot-cli" "templates.zip"
mv "$sbcli_archives_folder/spring-boot-cli" "$SBCLI_DIR/$SBCLI_VERSION"
unzip "$sbcli_archives_folder/templates.zip" -d "$SBCLI_DIR/$SBCLI_VERSION"
rm -f "$sbcli_archives_folder/templates.zip"
ln "$SBCLI_DIR/$SBCLI_VERSION/spring-boot-cli" "$SBCLI_DIR/current/sbcli"
