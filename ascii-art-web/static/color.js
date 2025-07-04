let globalColorValue = '#ff0000';
let targetColorValue = '#00ffff';
let backgroundColorValue = '#f8f9f9';

let globalPicker, targetPicker, backgroundPicker;

function initColors() {
  const pre = document.getElementById('asciiOutput');

  globalPicker = new iro.ColorPicker("#globalColorWheel", {
    width: 100,
    color: globalColorValue,
    borderWidth: 1,
    borderColor: "#ccc"
  });

  targetPicker = new iro.ColorPicker("#targetColorWheel", {
    width: 100,
    color: targetColorValue,
    borderWidth: 1,
    borderColor: "#ccc"
  });

  backgroundPicker = new iro.ColorPicker("#backgroundColorWheel", {
    width: 100,
    color: backgroundColorValue,
    borderWidth: 1,
    borderColor: "#ccc"
  });

  globalPicker.on('color:change', color => {
    globalColorValue = color.hexString;
    updateRadioSelection('color', globalColorValue);
    scheduleGenerate();
  });

  targetPicker.on('color:change', color => {
    targetColorValue = color.hexString;
    updateRadioSelection('targetColor', targetColorValue);
    scheduleGenerate();
  });

  backgroundPicker.on('color:change', color => {
    backgroundColorValue = color.hexString;
    updateRadioSelection('backgroundPreset', backgroundColorValue);
    pre.style.backgroundColor = backgroundColorValue;
  });

  document.querySelectorAll('input[name="color"]').forEach(r => {
    r.addEventListener('change', e => {
      globalColorValue = e.target.value;
      globalPicker.color.hexString = globalColorValue;
      scheduleGenerate();
    });
  });

  document.querySelectorAll('input[name="targetColor"]').forEach(r => {
    r.addEventListener('change', e => {
      targetColorValue = e.target.value;
      targetPicker.color.hexString = targetColorValue;
      scheduleGenerate();
    });
  });

  document.querySelectorAll('input[name="backgroundPreset"]').forEach(r => {
    r.addEventListener('change', e => {
      backgroundColorValue = e.target.value;
      backgroundPicker.color.hexString = backgroundColorValue;
      pre.style.backgroundColor = backgroundColorValue;
    });
  });
}

function updateRadioSelection(name, color) {
  const radios = document.querySelectorAll(`input[name="${name}"]`);
  radios.forEach(radio => {
    radio.checked = radio.value.toLowerCase() === color.toLowerCase();
  });
}
