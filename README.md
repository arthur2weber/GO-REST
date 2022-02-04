**Install**
----
  Just clone the repository in your local environment and run the command:
* **docker-compose up --build** 
  The endpoins are will be accessible at the address:
  * http://localhost:8080/
  * Ex: http://localhost:8080/fleets

  
**Run Tests**
  The tests are run in a separate container from the production container for this you must use the command:
* **docker-compose-f docker-compose.test.yml up --build --abort-on-container-exit**
  After running the tests, the container will be finalized


**Endpoints**


**Refrash Database**
----
  Erase all database data

* **URL**

  /database

* **Method:**

  `DELETE`
  
*  **URL Params**

   **Required:**

* **Data Params**

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
 
* **Error Response:**

  * **Code:** 304 Not Modified <br />
    **Content:** If there is a problem during Drop or Migrate



**List Fleets**
----
  List all Fleets

* **URL**

  /fleets

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**

* **Data Params**

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** [{ "id": int, "name": string, "max_speed": float64 }]
 
* **Error Response:**

  * **Code:** 304 Not Modified <br />
    **Content:** If there is a problem during Drop or Migrate



**Create Fleet**
----
  Create a Fleet

* **URL**

  /fleets

* **Method:**

  `POST`
  
*  **URL Params**
  * fleet_id: int

   **Required:**
    * name
    * max_speed

* **Data Params** { "name": string, "max_speed": float64 }

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** {"id": int}
 
* **Error Response:**

  * **Code:** 400 Bad Request <br />
    **Content:** Required Name

  OR

  * **Code:** 400 Bad Request <br />
  **Content:** Required Max_Speed
  If the max_ speed is empty or less than 0



**List FleetAlerts**
----
  List all fleet alerts of a fleet

* **URL**

  /fleets/{fleet_id}/alerts

* **Method:**

  `GET`
  
*  **URL Params**
  * fleet_id: int

   **Required:**

* **Data Params** 

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** [{"id": int, "fleet_id": int, "webhook": string }]

* **Error Response:**

  * **Code:** 400 Bad Request <br />
    **Content:** Invalid Fleet ID
    If the fllet id does not exist


**Create FleetAlerts**
----
  Create a Fleet Alert for a Fleet

* **URL**

  /fleets/{fleet_id}/alerts

* **Method:**

  `POST`
  
*  **URL Params**
  * fleet_id: int

   **Required:**
    * webhook
    * max_speed

* **Data Params** { "webhook": string, "max_speed": float64 }

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** {"id": int}
 
* **Error Response:**

  * **Code:** 400 Bad Request <br />
    **Content:** Invalid URL for Webhook
    If the webhook is not a valid url

  OR

  * **Code:** 400 Bad Request <br />
    **Content:** Invalid Fleet ID
    If the fllet id does not exist



**List Vehicles**
----
  List all vehicles

* **URL**

  /vehicles

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**

* **Data Params** 

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** [{"id": int, "name": string, "fleet_id": int, "max_speed": float64 }]
    If the vehicle does not have a maximum speed registered, the maximum speed of the fleet will be displayed by default. 

* **Error Response:**

  * **Code:** 400 Bad Request <br />
    **Content:** Invalid Fleet ID
    If the fllet id does not exist



**Create Vehicle**
----
  Create a Vehicle

* **URL**

  /vehicles

* **Method:**

  `POST`
  
*  **URL Params**
  * fleet_id: int

   **Required:**
    * name
    * fleet_id

* **Data Params** { "name": string, "fleet_id": int, "max_speed": float64 }

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** {"id": int}
 
* **Error Response:**

  * **Code:** 400 Bad Request <br />
    **Content:** Required Name
    If name value is enpty

  OR

  * **Code:** 400 Bad Request <br />
    **Content:** Required Name
    If name fleet_id is enpty or not found fleet



**List all Positions for a vehicle**
----
  List all Positions of a vehicle

* **URL**

  /vehicles/{vehicle_id}/positions

* **Method:**

  `GET`
  
*  **URL Params**
  * vehicle_id: int

   **Required:**

* **Data Params** 

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** [{"id": int, "vehicle_id": int, "timestamp": "ISO-8601", "latitude": float64, "longitude": float64, "max_speed": float64, "current_speed": float64}]

* **Error Response:**

  * **Code:** 400 Bad Request <br />
    **Content:** Invalid Vehicle ID
    If the vehicle id does not exist or vehicle not found


**Create o Positions for a vehicle**
----
  Create a Positions for a vehicle

* **URL**

  /vehicles/{vehicle_id}/possitions

* **Method:**

  `POST`
  
*  **URL Params**
  * vehicle_id: int

   **Required:**
    * timestamp
    * latitude
    * longitude
    * max_speed
    * current_speed

* **Data Params** {"timestamp": "ISO-8601", "latitude": float64, "longitude": float64, "max_speed": float64, "current_speed": float64}

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** {"id": int}
 
* **Error Response:**

  * **Code:** 400 Bad Request <br />
    **Content:** Invalid Vehicle ID
    If the vehicle id does not exist or vehicle not found

  OR

  * **Code:** 400 Bad Request <br />
    **Content:** Required timestamp
    If the timestamp value is empty

  OR

  * **Code:** 400 Bad Request <br />
    **Content:** Required latitude
    If the latitude value is empty

  OR

  * **Code:** 400 Bad Request <br />
    **Content:** Required longitude
    If the longitude value is empty

  OR

  * **Code:** 400 Bad Request <br />
    **Content:** Required current_speed
    If the current_speed value is empty

  After registration, if the vehicle is above its maximum speed (or the maximum speed of the fleet if the vehicle does not have it), its position will be sent to the webhooks registered in the vehicle fleet.