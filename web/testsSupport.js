function addCustomRow(tableName) {
    let table = document.getElementById(tableName).getElementsByTagName('tbody')[0];
    let id = table.getElementsByTagName("tr").length;
    let row = table.insertRow(id);

    //id cell
    let idCell = row.insertCell(0);
    idCell.innerText = id + 1;

    //question cell
    let questionInput = document.createElement("input");
    questionInput.type = "text";
    questionInput.placeholder = "Question text";
    questionInput.required = "true";
    let questionCell = row.insertCell(1);
    questionCell.appendChild(questionInput);


    let aInput = document.createElement("input");
    aInput.type = "text";
    aInput.placeholder = "Option a";
    aInput.required = "true";

    let bInput = document.createElement("input");
    bInput.type = "text";
    bInput.placeholder = "Option b";
    bInput.required = "true";

    let cInput = document.createElement("input");
    cInput.type = "text";
    cInput.placeholder = "Option c";
    cInput.required = "true";

    let dInput = document.createElement("input");
    dInput.type = "text";
    dInput.placeholder = "Option d";
    dInput.required = "true";


    let optionsDiv = document.createElement("div");
    let optionsCell = row.insertCell(2);
    optionsDiv.appendChild(aInput);
    optionsDiv.appendChild(bInput);
    optionsDiv.appendChild(cInput);
    optionsDiv.appendChild(dInput);
    optionsCell.appendChild(optionsDiv);


    let correctAnswerCell = row.insertCell(3);
    let selectDiv = document.createElement("select");
    selectDiv.classList.add("browser-default");

    let optionSelect = document.createElement("option");
    optionSelect.innerText = "Choose correct answer";
    optionSelect.setAttribute("disabled", "true");
    optionSelect.setAttribute("selected", "true");
    let optionA = document.createElement("option");
    optionA.value = "A";
    optionA.innerText = "A";
    let optionB = document.createElement("option");
    optionB.value = "B";
    optionB.innerText = "B";
    let optionC = document.createElement("option");
    optionC.value = "C";
    optionC.innerText = "C";
    let optionD = document.createElement("option");
    optionD.value = "D";
    optionD.innerText = "D";

    selectDiv.appendChild(optionSelect);
    selectDiv.appendChild(optionA);
    selectDiv.appendChild(optionB);
    selectDiv.appendChild(optionC);
    selectDiv.appendChild(optionD);
    selectDiv.id = 'label' + id;
    let divSelect = document.createElement("div");
    divSelect.classList.add("input-field");
    divSelect.appendChild(selectDiv);

    //input button
    let delButton = document.createElement("button");
    delButton.classList.add("waves-effect");
    delButton.classList.add("waves-light");
    delButton.classList.add("btn");
    delButton.addEventListener("click", function () {
        let rowId = this.closest('tr').rowIndex - 1;
        table.deleteRow(rowId);
        let rowsCounter = table.getElementsByTagName("tr").length;
        for (let i = 0; i < rowsCounter; i++) {
            let rowToUpdate = table.getElementsByTagName("tr")[i];
            rowToUpdate.getElementsByTagName("td")[0].innerHTML = i + 1;
        }
    });
    let iForButton = document.createElement("i");
    iForButton.classList.add("material-icons");
    iForButton.classList.add("left");
    iForButton.innerText = "delete";
    delButton.innerText = "Delete";
    delButton.appendChild(iForButton);
    let deleteCell = row.insertCell(4);
    deleteCell.style.textAlign = "center";
    deleteCell.appendChild(delButton);

    correctAnswerCell.appendChild(divSelect);
}

function createTestJSON() {
    let answers = '"answers":{"base":[';
    let BaseQuestionsTableData = '"baseQuestions":[';
    let rows = document.getElementById("baseQuestionsTable").getElementsByTagName("tbody")[0].getElementsByTagName("tr").length;
    for (let i = 0; i < rows; i++) {
        let row = document.getElementById("baseQuestionsTable").getElementsByTagName("tbody")[0].getElementsByTagName("tr")[i];
        let id = row.getElementsByTagName("td")[0].innerHTML;
        let question = row.getElementsByTagName("td")[1].getElementsByTagName("input")[0].value;
        let options = row.getElementsByTagName("td")[2].getElementsByTagName("input");
        let selectEl = row.getElementsByTagName("td")[3].getElementsByTagName("select")[0];
        let select = selectEl.options[selectEl.selectedIndex].innerText;
        BaseQuestionsTableData += '{';
        BaseQuestionsTableData += `"id" : ${id},`;
        BaseQuestionsTableData += `"question" : "${question}",`;
        BaseQuestionsTableData += `"optionA" : "${options[0].value}",`;
        BaseQuestionsTableData += `"optionB" : "${options[1].value}",`;
        BaseQuestionsTableData += `"optionC" : "${options[2].value}",`;
        BaseQuestionsTableData += `"optionD" : "${options[3].value}"`;
        BaseQuestionsTableData += "}";


        answers += `{"id": ${id}, "answer": "${select}"}`;
        if (i !== rows - 1) {
            BaseQuestionsTableData += ",";
            answers += ",";
        }
    }
    BaseQuestionsTableData += "],";
    answers += '],"reading" : [';

    let ReadingQuestionsTableData = `"reading": { "question" : "${document.getElementById("readingText").value}", `;
    ReadingQuestionsTableData += '"questions":[';
    let readingRows = document.getElementById("readingQuestionsTable").getElementsByTagName("tbody")[0].getElementsByTagName("tr").length;
    for (let i = 0; i < readingRows; i++) {
        let row = document.getElementById("readingQuestionsTable").getElementsByTagName("tbody")[0].getElementsByTagName("tr")[i];
        let id = row.getElementsByTagName("td")[0].innerHTML;
        let question = row.getElementsByTagName("td")[1].getElementsByTagName("input")[0].value;
        let options = row.getElementsByTagName("td")[2].getElementsByTagName("input");
        let selectEl = row.getElementsByTagName("td")[3].getElementsByTagName("select")[0];
        let select = selectEl.options[selectEl.selectedIndex].innerText;
        ReadingQuestionsTableData += '{';
        ReadingQuestionsTableData += `"id" : ${id},`;
        ReadingQuestionsTableData += `"question" : "${question}",`;
        ReadingQuestionsTableData += `"optionA" : "${options[0].value}",`;
        ReadingQuestionsTableData += `"optionB" : "${options[1].value}",`;
        ReadingQuestionsTableData += `"optionC" : "${options[2].value}",`;
        ReadingQuestionsTableData += `"optionD" : "${options[3].value}"`;
        ReadingQuestionsTableData += "}";

        answers += `{"id": ${id}, "answer": "${select}"}`;
        if (i !== readingRows - 1) {
            ReadingQuestionsTableData += ",";
            answers += ",";
        }
    }
    ReadingQuestionsTableData += "]},";
    answers += "]}";

    let writingTask = `"writing": "${document.getElementById("writingText").value}",`;

    let toSend = "{" + BaseQuestionsTableData + ReadingQuestionsTableData + writingTask + answers + "}";
    document.getElementById("hiddenHolder").value = toSend;
    console.log(toSend);
}
