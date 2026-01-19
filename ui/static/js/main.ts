

// rewrite using interface?!? maybe for more than 2 forms//

const form = document.getElementById("loginForm") as HTMLFormElement;

form.addEventListener("submit", (e) => {
    e.preventDefault();
    const usernameInput = document.getElementById("username") as HTMLInputElement;
    const username = usernameInput.value;

    const passwordInput = document.getElementById("password") as HTMLInputElement;
    const password = passwordInput.value;

    const nameError = document.getElementById("name-error") as HTMLSpanElement;
    nameError.textContent = "";

    const passwordError = document.getElementById("password-error") as HTMLSpanElement;
    passwordError.textContent = "";

    let noError = true; // TS Infer

    if (username === "" || username.length < 2) {
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

const createAccountForm = document.getElementById("createAccountForm") as HTMLFormElement;
createAccountForm.addEventListener("submit", (e) => {
    e.preventDefault();
    //sort out 409 conflict error ----- user already exists

    const usernameInput = document.getElementById("usernameCreate") as HTMLInputElement;
    const username = usernameInput.value;

    const passwordInput = document.getElementById("passwordCreate") as HTMLInputElement;
    const password = passwordInput.value;

    const nameError = document.getElementById("createusername-error") as HTMLSpanElement;
    nameError.textContent = "";

    const passwordError = document.getElementById("createpassword-error") as HTMLSpanElement;
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