# Design AsiaYo booking system
## search room
	1. client(search room info) request search service 
	2. search service get data(room info) from elasticsearch
	3. return dataset to client

## room reservation/edit order
	1. client(reserver room) request order service 
	2. order service(get room status) check room available, if unavailable return 
	3. create order, updata sql and elastic room available info and setting payment ttl
	4. send notification to operator
	5. return order and payment info to client

## cancel order
	1. disable specific order
	2. release vacancy to room available numbers
	3. update elasticsearch room info

## payment
	1. client request payment service 
	2. check payment ttl, if expired request order service to cancel order then return error to client
	3. if payment success, request order service(payment confirm)
	4. order service would update order psyment status
	5. createing payment record
	6. send notification to operator
	7. return payment status to client

## worker
#### clean up expired order
    1. activated by cronjob
    2. get expired order from redis
    3. request order service to cancel order
#### notice nearing expiration
    1. activated by cronjob
    2. get nearing expired order from redis
    3. sending alert to user