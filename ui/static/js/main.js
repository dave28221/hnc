document.addEventListener("DOMContentLoaded", () => {
    alert("hello");

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

        if (username === "") {
            nameError.textContent = "Please enter a valid name"
            noError = false;
        }

        if (password === "") {
            passwordError.textContent = "please enter a valid password"
            noError = false;
        }




        alert("Form submitted!");
    });
});
