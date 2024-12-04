uuid=`curl -X POST \
  http://localhost:8080/api/v1/receipts/process \
  -H 'content-type: application/json' \
  -d '{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}' | awk -F'id":"|["}]' '{print $3}'`
echo "Got back $uuid"

points=`curl -s "http://localhost:8080/api/v1/receipts/$uuid/points" | awk -F'points":|[}]' '{print $2}'`
echo "$uuid has $points points, should be 28"

uuid=`curl -X POST \
  http://localhost:8080/api/v1/receipts/process \
  -H 'content-type: application/json' \
  -d '{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}' | awk -F'id":"|["}]' '{print $3}'`
echo "Got back $uuid"

points=`curl -s "http://localhost:8080/api/v1/receipts/$uuid/points" | awk -F'points":|[}]' '{print $2}'`
echo "$uuid has $points points, should be 109"

