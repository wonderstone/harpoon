#!/bin/bash

# Check if yq is installed
if command -v yq &> /dev/null
then
    echo "yq is already installed"
else
    echo "yq is not installed. Installing now..."
    brew install yq    
    # Verify installation
    if command -v yq &> /dev/null
    then
        echo "yq has been successfully installed"
    else
        echo "Failed to install yq"
        exit 1
    fi
fi

# use yq to parse the config file info.yaml
# read all the values from the info.yaml file
# and store them in variables
DB_HOST=$(yq '.host' './dbscript/info.yaml')
DB_PORT=$(yq '.port' './dbscript/info.yaml')    
DB_USER=$(yq '.user' './dbscript/info.yaml')
DB_PASS=$(yq '.pass' './dbscript/info.yaml')
DB_NAME=$(yq '.name' './dbscript/info.yaml')

# create a new database
CONNECT_DB="mysql -h $DB_HOST -P $DB_PORT -u $DB_USER -p$DB_PASS -e"
CREATE_DB="CREATE DATABASE IF NOT EXISTS $DB_NAME;"
USE_DB="USE $DB_NAME;"


SQL_COMMAND="$CREATE_DB $USE_DB $CREATE_TABLE"
$CONNECT_DB "$SQL_COMMAND"
 
exit 0