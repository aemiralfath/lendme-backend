# final-project-backend

RESTFUL API of P2P Lending using Golang, PostgreSQL, and JWT

## Project Description

```markdown
    - assets              # Store any image or file
    - cmd                 # This is where the main.go located
    - config              # This is where we set up app configuration
    - dist                # This is where we set up swagger ui
    - internal            # This folder is used to store clean architecture folder
    - pkg                 # Utility Here
    - sql                 # sql scripts folder
```

## ERD

![len](assets/lendme.png)

## How to setup and run the code

### Step 1: Clean download libraries

from the root folder, run `go mod tidy`

### Step 2: Setup database and initial data

run sql script, from root folder you can check on `./sql/init.sql` (you can copy paste the script on DBEAVER)

### Step 3: Run the app

from the root folder, run `go run ./...`

### Step 4: API Doc

to see the api doc, you can open [http://localhost:8080/docs](http://localhost:8080/docs)

### Step 5: Check unit test coverage

from the root folder, run `make test-coverage`

## Note

list of transactions endpoints example:
- `localhost:8080/api/v1/transactions`
- `localhost:8080/api/v1/transactions?s=money&sortBy=amount&sort=desc&limit=10&page=1`
- `localhost:8080/api/v1/transactions?sortBy=recipient_id&sort=desc`
- `localhost:8080/api/v1/transactions?sortBy=amount&sort=asc`

