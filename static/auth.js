var xhr = new XMLHttpRequest();


function login_form() {
  let form = document.getElementById("form")
  form.innerHTML = ""

  let email_label = document.createElement("label")
  let email_input = document.createElement("input")
  email_input.setAttribute("id", "auth_email")
  email_input.setAttribute("type", "email")
  email_input.setAttribute("placeholder", "Введите email")
  email_label.innerHTML = "Логин"
  email_label.appendChild(email_input)

  let pass_label = document.createElement("label")
  let pass_input = document.createElement("input")
  pass_input.setAttribute("id", "auth_pass")
  pass_input.setAttribute("type", "password")
  pass_input.setAttribute("placeholder", "Введите пароль")
  pass_label.innerHTML = "Пароль"
  pass_label.appendChild(pass_input)

  form.appendChild(email_label)
  form.appendChild(pass_label)
  
  let submitBtn = document.createElement("button");
  submitBtn.innerHTML = "Войти";
  submitBtn.setAttribute("onclick", "login()");
  form.appendChild(submitBtn);

  let clearBtn = document.createElement("button");
  clearBtn.innerHTML = "Зарегистрироваться";
  clearBtn.setAttribute("onclick", "register_form()");
  form.appendChild(clearBtn);

}

function login() {
  let login = document.getElementById("auth_email").value
  let pass = document.getElementById("auth_pass").value
  console.log('login: ' + login + ' pass: ' + pass)
  let message = {
    email: login,
    password: pass
  }
  xhr.open("POST", "/auth")
  xhr.setRequestHeader("Content-type", "application/json; charset=utf-8");
  xhr.send(JSON.stringify(message))
  xhr.onreadystatechange = check;
}

//
function showError(message){
  let errDiv = document.createElement("div")
  let closeSpan = document.createElement("span") // document.getElementById("task-close") //
  closeSpan.addEventListener("click", function() {
    // errDiv.style.display = "none"
    // taskContainer.innerHTML = ""
    document.body.removeChild(errDiv)
  })
  closeSpan.setAttribute("class", "close")
  closeSpan.innerHTML = "&times"
  errDiv.setAttribute("class", "notification")
  errDiv.innerHTML = message
  errDiv.appendChild(closeSpan)
  // document.body.insertBefore(errDiv, document.body.lastChild)
  document.body.appendChild(errDiv)
}
//
  
function check() {
  if (xhr.readyState == 4 && xhr.status != 200){
    // alert(xhr.responseText)
    showError(xhr.responseText)
  } else if (xhr.readyState == 4 && xhr.status == 200){
    window.location.replace("/");
  }
}

function register_form() {
  let form = document.getElementById("form")
  form.innerHTML = ""

  let age_label = document.createElement("label")
  let age_input = document.createElement("input")
  age_input.setAttribute("id", "reg_age")
  age_input.setAttribute("type", "number")
  age_label.innerHTML = "Возраст"
  age_label.appendChild(age_input)

  let name_label = document.createElement("label")
  let name_input = document.createElement("input")
  name_input.setAttribute("id", "reg_name")
  name_label.innerHTML = "Имя"
  name_label.appendChild(name_input)

  let email_label = document.createElement("label")
  let email_input = document.createElement("input")
  email_input.setAttribute("id", "reg_email")
  email_input.setAttribute("type", "email")
  email_input.setAttribute("placeholder", "Введите email")
  email_label.innerHTML = "Email"
  email_label.appendChild(email_input)

  let pass_label = document.createElement("label")
  let pass_input = document.createElement("input")
  pass_input.setAttribute("id", "reg_pass")
  pass_input.setAttribute("type", "password")
  pass_input.setAttribute("placeholder", "Введите пароль")
  pass_label.innerHTML = "Пароль"
  pass_label.appendChild(pass_input)

  form.appendChild(name_label)
  form.appendChild(age_label)
  form.appendChild(email_label)
  form.appendChild(pass_label)
 
  let submitBtn = document.createElement("button");
  submitBtn.innerHTML = "Зарегистрироваться";
  submitBtn.setAttribute("onclick", "register()");
  form.appendChild(submitBtn);

  let clearBtn = document.createElement("button");
  clearBtn.innerHTML = "Войти";
  clearBtn.setAttribute("onclick", "login_form()");
  form.appendChild(clearBtn);

}

function register() {
  let user = {
    name: document.getElementById("reg_name").value,
    age: parseInt(document.getElementById("reg_age").value),
    email: document.getElementById("reg_email").value,
    password: document.getElementById("reg_pass").value
  }
  console.log("user: " + user.age)
  xhr.open("POST", "/register")
  xhr.setRequestHeader("Content-type", "application/json; charset=utf-8");
  xhr.send(JSON.stringify(user))
  xhr.onreadystatechange = check;


}



