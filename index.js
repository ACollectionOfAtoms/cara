let inputOne = document.querySelector('.file-input-one');
let inputTwo = document.querySelector('.file-input-two');
let previewOne = document.querySelector('.image-one');
let previewTwo = document.querySelector('.image-two');
let files = Object.seal([null, null]);

// input.style.opacity = 0;
inputOne.addEventListener('change', updateImageDisplay.bind(null, inputOne, previewOne));
inputTwo.addEventListener('change', updateImageDisplay.bind(null, inputTwo, previewTwo));

function updateImageDisplay(input, preview) {
  while(preview.firstChild) {
    preview.removeChild(preview.firstChild);
  }
  var curFiles = input.files;
  if(curFiles.length === 0) {
    preview.textContent = 'No files currently selected for upload';
  } else {
    const file = curFiles[0];
    if(validFileType(file)) {
      const url = window.URL.createObjectURL(file);
      preview.src = url;
    } else {
      preview.textContent = 'File name ' + file.name + ': Not a valid file type. Update your selection.';
    }
  }
}

const fileTypes = [
  'image/jpeg',
  'image/pjpeg',
  'image/png'
]

function validFileType(file) {
  return fileTypes.includes(file.type);
}
