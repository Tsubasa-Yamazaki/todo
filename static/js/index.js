function entry(id) {
	
	let imp = document.post.impName.value;
	let task = document.post.taskName.value;
	let limit = document.post.limitName.value;
	
	let table = document.getElementById(id);

	let newRow  = table.insertRow(-1);
  
	let cell1 = newRow.insertCell(-1);
	let cell2 = newRow.insertCell(-1);
	let cell3 = newRow.insertCell(-1);
	let cell4 = newRow.insertCell(-1);

	let impText  = document.createTextNode(imp);
	let taskText  = document.createTextNode(task);
	let limitText  = document.createTextNode(limit);

	cell1.appendChild(impText);
	cell2.appendChild(taskText);
	cell3.appendChild(limitText);
	cell4.innerHTML = '<button type="button" id="dButton" onclick="dButton(this)">削除</button>';
}

function dButton(obj) {
	tr = obj.parentNode.parentNode;
	tr.parentNode.deleteRow(tr.sectionRowIndex);
}