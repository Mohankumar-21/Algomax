function registerUser() {
    // Fetch API to send registration data to the server
    fetch('/api/add_user', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            name: document.getElementById('name').value,
            surname: document.getElementById('surname').value,
            username: document.getElementById('username').value,
            email: document.getElementById('email').value,
            password: document.getElementById('password').value,
        }),
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        alert(data.message);
    })
    .catch(error => console.error('Error:', error));
}

function loginUser() {
    // Fetch API to send login data to the server
    fetch('/api/sign_in', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            user: document.getElementById('loginUser').value,
            password: document.getElementById('loginPassword').value,
        }),
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        if (data.token) {
            alert('Login successful! Token: ' + data.token);
        } else {
            alert(data.error);
        }
    })
    .catch(error => console.error('Error:', error));
}