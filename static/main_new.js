var socket = new WebSocket("ws://localhost:8080/ws");
var usersTable = document.getElementById("users")
var usersTBody =  document.createElement("tbody")
usersTable.appendChild(usersTBody)
var tasksTable = document.getElementById("tasks")
var currentUser;

function getCookie(name) {
  const cookies = document.cookie.split(';');
  console.log('cookie len: ' + cookies.lenght + ' cookie: ' + cookies)
  for (let i = 0; i < cookies.length; i++) {
    const cookie = cookies[i].trim();
    if (cookie.startsWith(name + '=')) {
      return cookie.split('=')[1];
    }
  }
  return null;
}

function showError(message) {
  let errDiv = document.createElement("div")
  let closeSpan = document.createElement("span") // document.getElementById("task-close") //
  closeSpan.addEventListener("click", function() {
    document.body.removeChild(errDiv)
  })
  closeSpan.setAttribute("class", "close")
  closeSpan.innerHTML = "&times"
  errDiv.setAttribute("class", "notification")
  errDiv.innerHTML = message
  errDiv.appendChild(closeSpan)
  document.body.appendChild(errDiv)
}

console.log(usersTable)
socket.onopen = function(event) {
  console.log("Соединение открыто")

  const myCookieValue = getCookie("uuid_cookie");
  console.log(myCookieValue);

  var message = {
    type: "current user",
    obj: JSON.stringify(myCookieValue)
  }
  socket.send(JSON.stringify(message))

  message = {
    type: "find all users",
    obj: JSON.stringify("")
  }
  socket.send(JSON.stringify(message))
}


socket.onmessage = function(event) {
  console.log("we are in onmessage\n" + "event.data: " + event.data)
  var response = JSON.parse(event.data)
  if (response.error) {
    showError(response.obj)
    return
  }
  switch (response.type) {
    case "current user":
      setCurrentUser(response.obj)
      break
    case "find all users":
      FindAllUsers(response.obj)
      break
    case "create user":
      createUser(response.obj)
      break
    case "delete user":
      deleteUser(response.obj)
      break
    case "find user":
      findUser(response.obj)
      break
    case "update user by user":
      updateUserByUser(response.obj)
      break
    case "update user by admin":
      updateUserByAdmin(response.obj)
      break
    case "create task":
      createTask(response.obj)
      break
    case "delete task":
      deleteTask(response.obj)
      break
    case "update task":
      updateTask(response.obj)
      break
    case "find all task for user":
      findAllTaskForUser(response.obj)
      break
    case "create task for multiple users":
      createTaskForMultipleUsers(response.obj)
      break
    default:
      console.log("Ничего в свитче не подошло")
  }
}

socket.onclose = function(event) {
  console.log("Соединение закрыто")
}

function setCurrentUser(responseObj) {
  currentUser = JSON.parse(responseObj)
  console.log("set currenUser: " + currentUser)
  console.log("set currenUser string: " + currentUser.toString())
  if (currentUser.admin) {
    let createButton = document.createElement("button")
    createButton.innerHTML = "Создать пользователя"
    createButton.onclick = buttonUserCreate
    document.body.appendChild(createButton)

    let deleteButton = document.createElement("button")
    deleteButton.innerHTML = "Удалить пользователя"
    deleteButton.onclick = buttonUserDelete
    document.body.appendChild(deleteButton)
  }
}

function createUser(responseObj) {
  let user = JSON.parse(responseObj)
  addUser(user)
}

function addUser(user) {
  var userData = document.createElement("tr")
  userData.setAttribute("id", user.email)
  var email = document.createElement("td")
  var name = document.createElement("td")
  var age = document.createElement("td")
  email.innerHTML = user.email
  name.innerHTML = user.name
  age.innerHTML = user.age
  userData.appendChild(name)
  userData.appendChild(email)
  userData.appendChild(age)
  userData.addEventListener("click", function() {
    // window.history.pushState({}, '', user.email)
    window.location.hash = user.email
    var message = {
      type: "find all task for user",
      obj: JSON.stringify(user.email)
    }
    socket.send(JSON.stringify(message))
  })
  usersTBody.appendChild(userData)
  // usersTable.appendChild(userData)

}

function deleteUser(repsonseObj) {
  let deletedUser = document.getElementById(JSON.parse(repsonseObj).email)
  usersTBody.removeChild(deletedUser)
  // usersTable.removeChild(deletedUser)
}

function findUser(responseObj) { }

function updateUserByUser(responseObj) { }

function updateUserByAdmin(responseObj) { }

function createTask(responseObj) {
  let task = JSON.parse(responseObj)
  console.log("hash: ", window.location.hash)
  if (window.location.hash === ('#' + task.user_email)) {
    console.log("мы на нужном href: ", window.location.href)
    addTask(task)
  }
}

function addTask(item) {
  var task = document.createElement("tr")
  task.setAttribute("task_id", item.task_id)
  var task_id = document.createElement("td")     //uint64 `json:"task_id"`
  task_id.innerHTML = item.task_id
  var title = document.createElement("td")
  title.innerHTML = item.title //string `json:"title"`
  title.setAttribute("name", "title")
  var description = document.createElement("td")  //string `json:"description"`
  description.innerHTML = item.description
  description.setAttribute("name", "description")
  var status = document.createElement("td")    //string `json:"status"` // может лучше хранить числом 🤔
  status.innerHTML = item.status
  status.setAttribute("name", "status")
  var time_start = document.createElement("td")  //int64  `json:"time_start"`
  time_start.innerHTML = dateFormater(item.time_start)
  time_start.setAttribute("name", "time_start")
  var time_finish = document.createElement("td") //int64  `json:"time_finish"`
  time_finish.innerHTML = dateFormater(item.time_finish)
  time_finish.setAttribute("name", "time_finish")

  task.appendChild(title)
  task.appendChild(description)
  task.appendChild(status)
  task.appendChild(time_start)
  task.appendChild(time_finish)
  task.addEventListener('click', function() {
    fillModalBase()
    modalAddTaskTemplate(item)
    if (currentUser.admin || item.user_email === currentUser.email) {
      modalAddUpdateTaskButton()
      modalAddDeleteTaskButton()
    } else {
      offInput()
    }
  })
  tasksTable.appendChild(task)

}

function offInput() {
  document.querySelector('#my-container input[name="title"]').readOnly = true
  document.querySelector('#my-container input[name="description"]').readOnly = true
  document.querySelector('#my-container select[name="status"]').readOnly = true
  document.querySelector('#my-container input[name="time_start"]').readOnly = true
  document.querySelector('#my-container input[name="time_finish"]').readOnly = true
  let stopUser = document.createElement("label")
  stopUser.innerHTML = "Пользователь может редактировать только свои задания"
  stopUser.setAttribute("class", "stop-user")
  document.querySelector('#my-container').appendChild(stopUser)
}

function deleteTask(responseObj) {
  let task = document.querySelector('tr[task_id=' + CSS.escape(responseObj) + ']')
  if (task !== null) {
    tasksTable.removeChild(task)
  }

}

function updateTask(responseObj) {
  let task = JSON.parse(responseObj)
  if (window.location.hash === ('#' + task.user_email)) {
    let task_old = document.querySelector('tr[task_id=' + CSS.escape(task.task_id) + ']') //.getElementById(task.task_id)
    task_old.querySelector('td[name="title"]').innerHTML = task.title
    task_old.querySelector('td[name="description"]').innerHTML = task.description
    task_old.querySelector('td[name="status"]').innerHTML = task.status
    task_old.querySelector('td[name="time_start"]').innerHTML = dateFormater(task.time_start)
    task_old.querySelector('td[name="time_finish"]').innerHTML = dateFormater(task.time_finish)
  }

}

function findAllTaskForUser(responseObj) {
  var tasks = JSON.parse(responseObj)
  tasksTable.innerHTML = ""
  tasksTable.parentElement.removeChild(tasksTable)
  let userRow = document.getElementById(tasks[0].user_email)
  tasksTable = document.createElement("table")
  tasksTable.setAttribute("id", "tasks")
  userRow.appendChild(tasksTable)
  let header = document.createElement("tr")
  let title = document.createElement("th")
  title.innerHTML = "Название"
  let descr = document.createElement("th")
  descr.innerHTML = "Описание"
  let status = document.createElement("th")
  status.innerHTML = "Статус"
  let time_s = document.createElement("th")
  time_s.innerHTML = "Время начала"
  let time_f = document.createElement("th")
  time_f.innerHTML = "Время окончания"
  header.appendChild(title)
  header.appendChild(descr)
  header.appendChild(status)
  header.appendChild(time_s)
  header.appendChild(time_f)
  tasksTable.appendChild(header)
  tasks.forEach((item) => {
    addTask(item)
  })

}

function createTaskForMultipleUsers(responseObj) {
  let tasksAndErrors = JSON.parse(responseObj)
  tasksAndErrors.forEach((item) => {
    if (item.error !== "") {
      showError("Не смогли создать задание для " + item.task.user_email + ". Ошибка: " + item.error)
      return
    }
    createTask(JSON.stringify(item.task))
  })
}

function FindAllUsers(responseObj) {
  let header = document.createElement("tr")
  let name = document.createElement("th")
  name.innerHTML = "Имя"
  let email = document.createElement("th")
  email.innerHTML = "Email"
  let age = document.createElement("th")
  age.innerHTML = "Возраст"
  header.appendChild(name)
  header.appendChild(email)
  header.appendChild(age)
  // usersTable.appendChild(header)
  usersTBody.appendChild(header)
  var users = JSON.parse(responseObj)
  users.forEach((item) => {
    addUser(item)
  });
}

function buttonUserCreate() {
  fillModalBase()

  let email = document.createElement("input")
  email.setAttribute("name", "email")
  email.setAttribute("type", "email")
  email.setAttribute("pattern", "[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$")
  email.setAttribute("required", "");
  let email_label = document.createElement("label")
  email_label.innerHTML = "Email:"
  email_label.appendChild(email)
  let name = document.createElement("input")
  name.setAttribute("name", "name")
  let name_label = document.createElement("label")
  name_label.innerHTML = "Имя:"
  name_label.appendChild(name)
  let age = document.createElement("input")
  age.setAttribute("name", "age")
  age.setAttribute("type", "number")
  let age_label = document.createElement("label")
  age_label.innerHTML = "Возраст"
  age_label.appendChild(age)
  let password = document.createElement("input")
  password.setAttribute("name", "password")
  password.setAttribute("type", "password")
  let password_label = document.createElement("label")
  password_label.innerHTML = "Пароль:"
  password_label.appendChild(password)
  let admin = document.createElement("input")
  admin.setAttribute("name", "admin")
  admin.setAttribute("type", "checkbox")
  let admin_label = document.createElement("label")
  admin_label.innerHTML = "Администратор:"
  admin_label.appendChild(admin)

  let userContainer = document.querySelector('#my-container')
  userContainer.appendChild(name_label)
  userContainer.appendChild(age_label)
  userContainer.appendChild(admin_label)
  userContainer.appendChild(email_label)
  userContainer.appendChild(password_label)

  let buttonSave = document.createElement("button")
  buttonSave.innerHTML = "Сохранить"
  buttonSave.onclick = buttonSaveUser

  userContainer.appendChild(buttonSave)
}

function buttonSaveUser() {
  let name = document.querySelector('#my-container input[name="name"]').value
  let age = document.querySelector('#my-container input[name="age"]').value
  let admin = document.querySelector('#my-container input[name="admin"]').checked
  let email = document.querySelector('#my-container input[name="email"]').value
  let password = document.querySelector('#my-container input[name="password"]').value
  if(!email.match(
    /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
  )){
    showError("Вы ввели некорректный email")
    return
  }
  console.log('name: ' + name + '\nage: ' + age + '\nadmin: ' + admin + '\nemail: ' + email + '\npassword' + password)
  socket.send(JSON.stringify(
    {
      type: "create user",
      obj: JSON.stringify(
        {
          name: name,
          age: parseInt(age),
          admin: admin,
          email: email,
          password: password
        }
      )
    }
  ))
  offModal()
}

function offModal() {
  let modalDiv = document.getElementById("my-modal")
  modalDiv.style.display = "block"
  let modalContent = document.getElementById("my-modal-content")
  let userContainer = document.getElementById("my-container")
  let closeSpan = document.querySelector("#my-modal-content span")
  modalContent.removeChild(closeSpan)
  modalDiv.style.display = "none"
  userContainer.innerHTML = ""

}

function buttonUserDelete() {
  fillModalBase()
  modalAddCheckbox()
  let taskContainer = document.getElementById("my-container")
  let deleteButton = document.createElement("button")
  deleteButton.innerHTML = "Удалить выбранных пользователей"
  deleteButton.onclick = actionForDeleteButton
  taskContainer.appendChild(deleteButton)
}

function actionForDeleteButton() {
  var newTable = document.getElementById('new-table');
  var selectedEmails = [];
  newTable.querySelectorAll('input[type="checkbox"]:checked').forEach(function(checkbox) {
    var surname = checkbox.parentNode.previousSibling.textContent;
    selectedEmails.push(surname);
  });
  console.log("Emails: " + selectedEmails)
  if (selectedEmails.length == 0) {
    showError("Не выбран ни один пользователь")
    return
  }

  offModal()

  selectedEmails.forEach((email) =>
    socket.send(JSON.stringify({
      type: "delete user",
      obj: JSON.stringify({
        email: email
      })
    }))
  )


}

function buttonTaskCreate() {
  let task = {
    task_id: "",
    title: "",
    description: "",
    status: "",
    time_start: parseInt(Date.now() / 100),
    time_finish: parseInt(Date.now() / 100),
    user_email: ""
  }
  fillModalBase()
  if (!currentUser.admin) {
    task.user_email = currentUser.email
    modalAddTaskTemplate(task)
    buttonCreateTaskByUser()
    return
  }
  modalAddTaskTemplate(task)
  modalAddCheckbox()
  buttonCreateTasksByAdmin()
}

function fillModalBase() {
  let modalDiv = document.getElementById("my-modal") //document.createElement("div")
  modalDiv.style.display = "block"
  let modalContent = document.getElementById("my-modal-content") //document.createElement("div")
  let taskContainer = document.getElementById("my-container")
  let closeSpan = document.createElement("span") // document.getElementById("task-close") //
  closeSpan.addEventListener("click", function() {
    modalDiv.style.display = "none"
    taskContainer.innerHTML = ""
    modalContent.removeChild(closeSpan)
  })
  closeSpan.setAttribute("class", "close")
  closeSpan.innerHTML = "&times"

  modalContent.appendChild(taskContainer)
  modalContent.insertBefore(closeSpan, taskContainer)
  modalDiv.appendChild(modalContent)
}

function modalAddTaskTemplate(task) {

  let taskContainer = document.getElementById("my-container")
  var task_id = document.createElement("div")
  task_id.setAttribute("type", "hidden")
  task_id.setAttribute("task_id", task.task_id)
  task_id.setAttribute("name", "task_id")
  var title = document.createElement("input")
  title.setAttribute("value", task.title) //string `json:"title"`
  title.setAttribute("name", "title")
  let title_label = document.createElement("label")
  title_label.innerHTML = "Название:"
  title_label.appendChild(title)
  var description = document.createElement("input")  //string `json:"description"`
  description.setAttribute("value", task.description)
  description.setAttribute("name", "description")
  let description_label = document.createElement("label")
  description_label.innerHTML = "Описание:"
  description_label.appendChild(description)
  var status = document.createElement("select")    //string `json:"status"` // может лучше хранить числом 🤔
  status.setAttribute("value", task.status)
  status.setAttribute("name", "status")
  let pending = document.createElement("option")
  pending.innerHTML = "pending"
  let completed = document.createElement("option")
  completed.innerHTML = "completed"
  let deleted = document.createElement("option")
  deleted.innerHTML = "deleted"
  let overdue = document.createElement("option")
  overdue.innerHTML = "overdue"
  status.appendChild(pending)
  status.appendChild(completed)
  status.appendChild(overdue)
  status.appendChild(deleted)
  let status_label = document.createElement("label")
  status_label.innerHTML = "Статус:"
  status_label.appendChild(status)
  var time_start = document.createElement("input")  //int64  `json:"time_start"`
  time_start.setAttribute("value", dateFormater(task.time_start))
  // time_start.innerHTML = dateFormater(task.time_start)
  time_start.setAttribute("type", "date")
  time_start.setAttribute("name", "time_start")
  let time_start_label = document.createElement("label")
  time_start_label.innerHTML = "Время начала:"
  time_start_label.appendChild(time_start)
  var time_finish = document.createElement("input") //int64  `json:"time_finish"`
  time_finish.setAttribute("value", dateFormater(task.time_finish))
  // time_finish.innerHTML = dateFormater(task.time_finish)
  time_finish.setAttribute("type", "date")
  time_finish.setAttribute("name", "time_finish")
  let time_finish_label = document.createElement("label")
  time_finish_label.innerHTML = "Время окончания:"
  time_finish_label.appendChild(time_finish)

  taskContainer.appendChild(task_id)
  taskContainer.appendChild(title_label)
  taskContainer.appendChild(description_label)
  taskContainer.appendChild(status_label)
  taskContainer.appendChild(time_start_label)
  taskContainer.appendChild(time_finish_label)

}


function modalAddCheckbox() {
  let modalDiv = document.getElementById("my-modal")
  modalDiv.style.display = "block"
  let taskContainer = document.getElementById("my-container")

  var newTable = document.createElement('table');
  newTable.setAttribute("id", "new-table")
  let thead = document.createElement("thead")
  var headerRow = document.createElement('tr');
  var nameHeader = document.createElement('th');
  nameHeader.textContent = 'Имя';
  headerRow.appendChild(nameHeader);
  var emailHeader = document.createElement('th');
  emailHeader.textContent = 'Email';
  headerRow.appendChild(emailHeader);
  var checkboxHeader = document.createElement('th');
  checkboxHeader.textContent = 'Выбрать';
  headerRow.appendChild(checkboxHeader);
  thead.appendChild(headerRow)
  newTable.appendChild(thead);

  usersTable.querySelectorAll('tr[id]').forEach(function(row, index) {
    if (index > 0) {
      var name = row.querySelectorAll('td')[0].textContent;
      var email = row.querySelectorAll('td')[1].textContent;
      var checkboxCell = document.createElement('td');
      var checkbox = document.createElement('input');
      checkbox.type = 'checkbox';
      checkbox.name = 'user';
      checkbox.value = name + ' ' + email;
      checkboxCell.appendChild(checkbox);
      var newRow = document.createElement('tr');
      var nameCell = document.createElement('td');
      nameCell.textContent = name;
      newRow.appendChild(nameCell);
      var emailCell = document.createElement('td');
      emailCell.textContent = email;
      newRow.appendChild(emailCell);
      newRow.appendChild(checkboxCell);
      newTable.appendChild(newRow);
    }
  });

  taskContainer.appendChild(newTable)

}

function buttonCreateTasksByAdmin() {
  let taskContainer = document.getElementById("my-container")
  let createButton = document.createElement("button")
  createButton.innerHTML = "Выдать задание"
  createButton.onclick = actionButtonCreateTasksByAdmin
  taskContainer.appendChild(createButton)

}


function actionButtonCreateTasksByAdmin() {
  var newTable = document.getElementById('new-table');
  var selectedEmails = [];
  newTable.querySelectorAll('input[type="checkbox"]:checked').forEach(function(checkbox) {
    var surname = checkbox.parentNode.previousSibling.textContent;
    selectedEmails.push(surname);
  });
  console.log("Emails: " + selectedEmails)
  if (selectedEmails.length == 0) {
    showError("Не выбран ни один пользователь")
    return
  }

  let task = taskSelector()

  offModal()

  socket.send(JSON.stringify({
    type: "create task for multiple users",
    obj: JSON.stringify({
      task: task,
      emails: selectedEmails
    })
  }))
}

function taskSelector() {
  let title = document.querySelector('#my-container input[name="title"]').value
  let description = document.querySelector('#my-container input[name="description"]').value
  let status = document.querySelector('#my-container select[name="status"]').value
  let time_start = document.querySelector('#my-container input[name="time_start"]').value
  let time_finish = document.querySelector('#my-container input[name="time_finish"]').value

  let task = {
    title: title,
    description: description,
    status: status,
    time_start: Math.floor(new Date(time_start).getTime() / 1000),
    time_finish: Math.floor(new Date(time_finish).getTime() / 1000),
  }
  return task
}

function buttonCreateTaskByUser() {
  let taskContainer = document.getElementById("my-container")
  let createButton = document.createElement("button")
  createButton.innerHTML = "Выдать задание"
  createButton.onclick = actionButtonCreateTasksByUser
  taskContainer.appendChild(createButton)
}

function actionButtonCreateTasksByUser() {
  let task = taskSelector()

  socket.send(JSON.stringify(
    {
      type: "create task",
      obj: JSON.stringify(
        {
          title: task.title,
          description: task.description,
          status: task.status,
          time_start: task.time_start,
          time_finish: task.time_finish,
          user_email: currentUser.email
        }

      )
    }))
  offModal()

}

function modalAddUpdateTaskButton() {
  let updateButton = document.createElement("button")
  updateButton.innerHTML = "Обновить данные"
  updateButton.onclick = actionUpdateTaskButton
  document.querySelector('#my-container').appendChild(updateButton)
}

function actionUpdateTaskButton() {
  let task = taskSelector()
  let task_id = document.querySelector('#my-container div[name="task_id"]').getAttribute("task_id")
  console.log('task id: ' + task_id + ' to int: ' + parseInt(task_id))
  socket.send(JSON.stringify(
    {
      type: "update task",
      obj: JSON.stringify(
        {
          task_id: parseInt(task_id),
          title: task.title,
          description: task.description,
          status: task.status,
          time_start: task.time_start,
          time_finish: task.time_finish,
          user_email: window.location.hash.toString().substring(1) // pathname.toString().substring(1)
        }

      )
    }))
  offModal()
}

function modalAddDeleteTaskButton() {
  let deleteButton = document.createElement("button")
  deleteButton.innerHTML = "Удалить"
  deleteButton.onclick = actionDeleteTaskButton
  document.querySelector('#my-container').appendChild(deleteButton)
}

function actionDeleteTaskButton() {
  let task_id = document.querySelector('#my-container div[name="task_id"]').getAttribute("task_id")
  socket.send(JSON.stringify(
    {
      type: "delete task",
      obj: JSON.stringify(parseInt(task_id))
    }
  ))
  offModal()
}

function dateFormater(unixTime) {
  var date = new Date(unixTime * 1000);
  return date.toISOString().substring(0, 10);
}

function logout(){
  // window.location.href = "/logout"
  window.location.replace("/login")
}
