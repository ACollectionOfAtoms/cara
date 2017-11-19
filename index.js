let inputOne = document.querySelector('.file-input-one');
let inputTwo = document.querySelector('.file-input-two');
let previewOne = document.querySelector('.image-one');
let previewTwo = document.querySelector('.image-two');
let files = Object.seal([null, null]);

// input.style.opacity = 0;
inputOne.addEventListener('change', updateImageDisplay.bind(null, inputOne, previewOne));
inputTwo.addEventListener('change', updateImageDisplay.bind(null, inputTwo, previewTwo));
uploadedFiles = {};

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

    uploadedFiles[fileName] = file;

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

function uploadFiles() {
  const fileCount = Object.keys(uploadedFiles).length;
  if (fileCount !== 2) {
    window.alert('please include a total of two faces');
    return;
  }
  const uri = '/upload';
  const formData = new FormData();
  const files = Object.values(uploadedFiles);
  console.log(files);
  files.forEach((file, index)=> {
    formData.append(`file-${index}`, file);
  });
  for (var par of formData.entries()) {
    console.log(par[0] + ',' + par[1]);
  }

  fetch(uri, {
    method: 'POST',
    'Content-Type': 'multipart/form-data',
    body: formData
  }).then(res => {
    console.log(res);
  }).catch(e => {
    console.log('Error', e);
  })
}

const form = document.getElementById('upload-form');
form.addEventListener('submit', e => {
  e.preventDefault();
  uploadFiles();
})
