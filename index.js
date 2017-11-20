let inputOne = document.querySelector('.file-input-one');
let inputTwo = document.querySelector('.file-input-two');
let previewOne = document.querySelector('.image-one');
let previewTwo = document.querySelector('.image-two');

inputOne.addEventListener('change', updateImageDisplay.bind(null, inputOne, previewOne));
inputTwo.addEventListener('change', updateImageDisplay.bind(null, inputTwo, previewTwo));
const uploadedFiles = {};

function updateImageDisplay(input, preview) {
  while(preview.firstChild) {
    preview.removeChild(preview.firstChild);
  }
  const fileName = preview.className;
  var curFiles = input.files;
  if(curFiles.length === 0) {
    preview.textContent = 'No files currently selected for upload';
  } else {
    const file = curFiles[0];
    if(validFileType(file)) {
      const url = window.URL.createObjectURL(file);
      preview.src = url;
      // store the file to be uploaded...
      uploadedFiles[fileName] = file;
    } else {
      preview.textContent = 'File name ' + file.name + ': Not a valid file type. Update your selection.';
    }
  }
}

function validFileType(file) {
  const fileTypes = [
    'image/jpeg',
    'image/pjpeg',
    'image/png'
  ]
  return fileTypes.includes(file.type);
}

function uploadFiles() {
  const fileCount = Object.keys(uploadedFiles).length;
  if (fileCount !== 2) {
    window.alert('please include a total of two faces');
    return;
  }
  const uri = '/upload';
  const formData = new FormData();
  const files = Object.values(uploadedFiles);
  const resultDiv = document.querySelector('#result');

  console.log(files);
  files.forEach((file, index)=> {
    formData.append(`file-${index}`, file);
  });
  resultDiv.textContent = 'Comparing...'
  fetch(uri, {
    method: 'POST',
    'Content-Type': 'multipart/form-data',
    body: formData
  })
  .then(r => r.text())
  .then(data => {
    let res = Math.round(data, 2)
    res = `${res}% Similar!`
    resultDiv.textContent = res;
  }).catch(e => {
    console.log('Error', e);
  })
}

const form = document.getElementById('upload-form');
form.addEventListener('submit', e => {
  e.preventDefault();
  uploadFiles();
})
