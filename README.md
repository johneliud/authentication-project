# Authentication Project

Welcome to the Authentication Project! This project provides a robust and secure authentication system, including user registration, email verification, and sign-in functionalities. It's designed to be a foundational component for web applications requiring user authentication.

## Features

- **User Registration**: Allows users to create an account with their personal details and password.
- **Email Verification**: Sends a verification code to the user's email address to confirm their identity.

## Installation

To set up the project locally, follow these steps:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/johneliud/authentication-project.git
   ```
2. **Navigate to the Project Directory**:
   ```bash
   cd authentication-project
   ```
3. **Set Up Environment Variables**:
   - Create a `.env` file in the root directory.
   - Add the following environment variables:
     ```env
     SMTP_EMAIL=your-email@example.com
     SMTP_PASSWORD=your-email-password
     SMTP_SERVER=smtp.gmail.com
     SMTP_PORT=587
     ```

_Note: Replace placeholder values like `your-email@example.com` and `your-username` with your actual information when setting up the project._

4. **Run the Application**:
   ```bash
   go run .
   ```

## Usage

- **Sign-Up**: Users can register by providing their first name, last name, email, and password. Upon successful registration, a verification code is sent to their email.
- **Email Verification**: After registration, users should enter the verification code received via email to verify their account.
- **Sign-In**: Once verified, users can sign in using their email and password.

## Contributing

We welcome contributions to enhance the Authentication Project. To contribute, please follow these guidelines:

1. **Fork the Repository**: Click on the 'Fork' button at the top right corner of the repository page to create a copy of the repository in your GitHub account.
2. **Clone Your Fork**:
   ```bash
   git clone https://github.com/your-username/authentication-project.git
   ```
3. **Create a New Branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```
4. **Make Your Changes**: Implement your feature or fix the identified issue.
5. **Commit Your Changes**:
   ```bash
   git commit -m "Add your commit message here"
   ```
6. **Push to Your Fork**:
   ```bash
   git push origin feature/your-feature-name
   ```
7. **Submit a Pull Request**: Navigate to the original repository and click on 'New Pull Request'. Provide a clear description of your changes and submit the pull request for review.

Please ensure your contributions align with the project's coding standards and include appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/johneliud/authentication-project/blob/main/LICENSE) file for more details.

## Contact

For any questions or suggestions, please contact [me](mailto:johneliud4@gmail.com).
