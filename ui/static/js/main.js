document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("loginForm");

    form.addEventListener("submit", (e) => {
        e.preventDefault();
        const username = document.getElementById("username").value;
        const password = document.getElementById("password").value;

        const nameError = document.getElementById("name-error");
        const passwordError = document.getElementById("password-error");
        nameError.textContent = "";
        passwordError.textContent = "";

        let noError = true;

        if (username === "" && username.length < 2) {
            nameError.textContent = "Please enter a valid name longer that is longer 2 characters"
            noError = false;
        }

        if (password === "") {
            passwordError.textContent = "please enter a valid password"
            noError = false;
        }
        if (!noError) {
            e.preventDefault();
        }


    });


    const createAccountForm = document.getElementById("createAccountForm");

    createAccountForm.addEventListener("submit", (e) => {
        const username = document.getElementById("usernameCreate").value;
        const password = document.getElementById("passwordCreate").value;

        const nameError = document.getElementById("createusername-error");
        const passwordError = document.getElementById("createpassword-error");
        nameError.textContent = "";
        passwordError.textContent = "";

        let noError = true;

        if (username === "" || username.length < 2 || username.length > 15) {
            nameError.textContent = "Username must be between 2 and 15 characters"
            noError = false;
        }

        if (password === "") {
            passwordError.textContent = "Please enter a password"
            noError = false;
        }

        if (!noError) {
            e.preventDefault();
        }

    });
});
