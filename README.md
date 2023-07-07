# razorpay-cli

## Fetch customers

### go run main.go customer

## Create payment

### go run main.go payment create --amount=112 --method="upi" --vpa="manask22221@okicici"

## Fetch payment

### go run main.go payment fetch --id="pay_MAVhcpLPpG00kd"

## Capture payment

### go run main.go payment capture --payment_id="pay_MAVhcpLPpG00kd" --amount=100

## Order create

### go run main.go order create --amount=100 --receipt="Receipt no. 1"

## Order fetch

### go run main.go order fetch --id="order_MAcMQWQ5eQoFiC"