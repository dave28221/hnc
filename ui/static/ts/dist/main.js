"use strict";
document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById("loginForm");
    form.addEventListener("submit", (e) => {
        const usernameInput = document.getElementById("username");
        const username = usernameInput.value;
        const passwordInput = document.getElementById("password");
        const password = passwordInput.value;
        const nameError = document.getElementById("name-error");
        nameError.textContent = "";
        const passwordError = document.getElementById("password-error");
        passwordError.textContent = "";
        let noError = true; // TS Infer
        if (username === "" || username.length < 2) {
            nameError.textContent = "Please enter a valid name longer that is longer 2 characters";
            noError = false;
        }
        if (password === "") {
            passwordError.textContent = "please enter a valid password";
            noError = false;
        }
        if (!noError) {
            e.preventDefault();
        }
    });
    const createAccountForm = document.getElementById("createAccountForm");
    createAccountForm.addEventListener("submit", (e) => {
        //sort out 409 conflict error ----- user already exists
        const usernameInput = document.getElementById("usernameCreate");
        const username = usernameInput.value;
        const passwordInput = document.getElementById("passwordCreate");
        const password = passwordInput.value;
        const nameError = document.getElementById("createusername-error");
        nameError.textContent = "";
        const passwordError = document.getElementById("createpassword-error");
        passwordError.textContent = "";
        let noError = true;
        if (username === "" || username.length < 2 || username.length > 15) {
            nameError.textContent = "Username must be between 2 and 15 characters";
            noError = false;
        }
        if (password === "") {
            passwordError.textContent = "Please enter a password";
            noError = false;
        }
        if (!noError) {
            e.preventDefault();
        }
    });
});
// create 409 conflict error
//# sourceMappingURL=main.js.map