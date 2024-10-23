// app.js

document.getElementById("goEditorPageButton").onclick = function () {
    location.href = "editorpage.html";
};


// Deal with input image

// Connect clickable text to html input selection
document.getElementById("uploadInputImageButton").onclick = function () {
    document.getElementById('inputImage').click();
};

// Handle output of input selection
document.getElementById('inputImage').addEventListener('change', function(event) {
    const file = event.target.files[0];

    if (file) {
	if (file.size > 25 * 1024 * 1024) {
	    alert('File size exceeds 25MB. Please select a smaller file.');
	    this.value = '';
	} else {
	    console.log('File selected:', file.name);

	    // Display image preview
	    const reader = new FileReader();
            reader.onload = function (e) {
                previewImage.src = e.target.result;
                previewImage.style.display = 'block';
            };
            reader.readAsDataURL(file);
	}
    }
});

