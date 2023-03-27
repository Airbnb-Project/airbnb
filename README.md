# AIRBNB Project

Airbnb project is an application used to find places for stay. This project was developed using GO based on Clean Architecture.

<div>
<h3>Features</h3>
<details><summary>👤User</summary>
<p>
User can register and login as a "user" and "host"

| Method | Endpoint  | JWT Token | Role  | Function                                    |
| ------ | --------- | --------- | ----- | ------------------------------------------- |
| POST   | /register | -         | -     | this is how users register their account.   |
| POST   | /login    | NO        | -     | this is how users login.                    |
| GET    | /users    | YES       | user  | user obtain their account information.      |
| PUT    | /users    | YES       | user  | this is how users update their profile.     |
| DELETE | /users    | YES       | user  | this is how users deactive their profile.   |

</p>
</details>

<details><summary>🏡Homestay</summary>
<p>
Host can add a homestay and manage the homestay for user.

| Method | Endpoint                      | JWT Token | Role                 | Function                                                                                                                                               |
| ------ | -------------------------     | --------- | -------------------- | ----------------------------------------------- |
| POST   | /homestays                    | YES       | host                 | host can add places for homestay.               |
| GET    | /homestays?page={page_number} | NO        | user/host            | user/host can see all the places for homestay.  |
| GET    | /homestays/{id}               | NO        | user/host            | user/host can see homestay detail by id.        |
| GET    | /myhomestay                   | YES       | host                 | host can see their homestay.                    |
| PUT    | /homestays/{id}               | YES       | host                 | host can edit/update their homestay.            |
| DELETE | /homestays                    | YES       | host                 | host can delete their homestay.                 |

</p>
</details>

<details><summary>🌟Feedback</summary>
<p>
User can gives feedback on homestay that users stay.

| Method | Endpoint     | JWT Token | Role      | Function                                     |
| ------ | ------------ | --------- | --------- | -------------------------------------------- |
| POST   | /feedbacks   | YES       | user      | user can gives a feedback for the homestay.  |
| GET    | /feedbacks   | NO        | user/host | user/host can see all feedbacks.             |
| GET    | /myfeedback  | YES       | user      | user can see their feedback.                 |

</p>
</details>

<details><summary>🎫Reservation</summary>
<p>
User can reserve a homestay.

| Method | Endpoint                  | JWT Token | Role      | Function                                          |
| ------ | ------------------------- | --------- | --------- | ------------------------------------------------- |
| POST   | /reservations             | YES       | user      | user can reserve a homestay.                      |
| GET    | /reservations             | YES       | user      | user can see their own reservation.               |
| GET    | /reservations/{id}        | YES       | user      | user can see the detail of their own reservation. |
| PUT    | /reservations/{id}/accept | YES       | host      | host can manual accept the reservation.           |
| PUT    | /reservations/{id}/cancel | YES       | user/host | user/host can cancel the reservation.             |
| POST   | /callback                 | YES       | user      | retrieving transaction data from midtrans.        |

</p>
</details>

</div>

## Tech Stack

## ERD

## Unit Test

## Run Locally

1. Clone the project

    ```bash
    $ git clone https://github.com/Airbnb-Project/airbnb.git
    ```

2. Go to project directory

    ```bash
    $ cd airbnb
    ```

3. Dont forget to activate the credential for third party api like cloudinary and midtrans
4. Create local.env like on local.env.example
5. Download all packages and dependencies
    ```bash
    $ go mod tidy
    ```
6. Run the program
    ```bash
    $ go run .
    ```
7. Enjoy

## Authors
[![GitHub](https://img.shields.io/badge/griffin-%23121011.svg?style=for-the-badge&logo=github&logoColor=blue)](https://github.com/kgriffinh)
