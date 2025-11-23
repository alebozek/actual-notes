function passwordChecks() {
    const pass_regex = /^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[@#$%^&+=-_])(?=\S+$).{8,}$/;
    const error_div = document.getElementById("error-msg");
    const submit_btn = document.getElementById("submit_btn");

    const pass = document.getElementById("password").value;
    const confPass = document.getElementById("confirm_password").value;


    if (pass !== confPass) {
        error_div.innerHTML = "Passwords must be the same.";
        error_div.style.display = 'block';
        submit_btn.style.backgroundColor = '#455045ff';
        submit_btn.disabled = true;
    } 
    else if (!pass.match(pass_regex)) {
        error_div.innerHTML = "The password must contain at least one digit, one lowercase character, an uppercase character, a special character, and at least 8 characters.";
        error_div.style.display = 'block';
        submit_btn.style.backgroundColor = '#455045ff';
        submit_btn.disabled = true;
    } 
    else {
        error_div.style.display = 'none';
        submit_btn.disabled = false;
        submit_btn.style.backgroundColor = '#4CAF50';  
    }
}

function loginPasswordChecks() {
    const pass_regex = /^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[@#$%^&+=-_])(?=\S+$).{8,}$/;
    const error_div = document.getElementById("error-msg");
    const submit_btn = document.getElementById("submit_btn");

    const pass = document.getElementById("password").value;

    if (!pass.match(pass_regex)) {
        error_div.innerHTML = "The password must contain at least one digit, one lowercase character, an uppercase character, a special character, and at least 8 characters.";
        error_div.style.display = 'block';
        submit_btn.style.backgroundColor = '#455045ff';
        submit_btn.disabled = true;
    } 
    else {
        error_div.style.display = 'none';
        submit_btn.disabled = false;
        submit_btn.style.backgroundColor = '#4CAF50';  
    }
}