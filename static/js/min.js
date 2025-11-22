function passwordChecks() {
    pass_regex = "^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[@#$%^&+=])(?=\S+$).{8,}$"
    error_div = document.getElementById("error-msg");
    submit_btn = document.getElementById("submit_btn");

    pass = document.getElementById("password").value;
    confPass = document.getElementById("confirm_password").value;

    if (pass != confPass || pass.match(pass_regex)){
        error_div.innerHTML = "The password must contain at least one digit, one lowercase character, an uppercase character, a special character and at least 8 characters.";
        error_div.style.display='block';
        submit_btn.style.backgroundColor = '#455045ff';
        submit_btn.disabled = true;
    }else {
        error_div.style.display = 'none';
        submit_btn.disabled = false;
        submit_btn.style.backgroundColor = '#4CAF50';
    }
}