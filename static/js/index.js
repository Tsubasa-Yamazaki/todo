

function entry(id) {
	
	let imp = document.post.impName.value;
	let task = document.post.taskName.value;
	let limit = document.post.limitName.value;
	
	let table = document.getElementById(id);

	let newRow   = table.insertRow(-1);
  
	let cell1 = newRow.insertCell(-1);
	let cell2 = newRow.insertCell(-1);
	let cell3 = newRow.insertCell(-1);
	let newText1  = document.createTextNode(imp);
	let newText2  = document.createTextNode(task);
	let newText3  = document.createTextNode(limit);

	cell1.appendChild(newText1);
	cell2.appendChild(newText2);
	cell3.appendChild(newText3);
}