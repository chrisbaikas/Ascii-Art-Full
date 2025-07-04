let debounceTimer;
const DEBOUNCE_MS = 500;
let justCleared = false;

function initGenerate() {
  const form = document.getElementById('asciiForm');
  const textarea = document.getElementById('inputText');
  const counter = document.getElementById('charCounter');
  const clearBtn = document.getElementById('clearBtn');
  const pre = document.getElementById('asciiOutput');
  const err = document.getElementById('error');

  textarea.addEventListener('input', () => {
    const len = textarea.value.length;
    counter.textContent = `${len}/1000000`;
    counter.classList.toggle('over-limit', len > 1000000);
  });

  clearBtn.addEventListener('click', () => {
    form.reset();
    counter.textContent = '0/1000000';
    pre.textContent = '';
    pre.className = 'align-left';
    err.hidden = true;
    err.innerHTML = '';
    justCleared = true;

    globalColorValue = '#ff0000';
    targetColorValue = '#00ffff';
    backgroundColorValue = '#f8f9f9';

    globalPicker.color.hexString = globalColorValue;
    targetPicker.color.hexString = targetColorValue;
    backgroundPicker.color.hexString = backgroundColorValue;
    pre.style.backgroundColor = backgroundColorValue;

    updateRadioSelection('color', globalColorValue);
    updateRadioSelection('targetColor', targetColorValue);
    updateRadioSelection('backgroundPreset', backgroundColorValue);
  });

  form.addEventListener('input', scheduleGenerate);
  form.addEventListener('change', scheduleGenerate);
  form.addEventListener('submit', e => {
    e.preventDefault();
    doGenerate();
  });
}

function scheduleGenerate() {
  if (justCleared) {
    justCleared = false;
    return;
  }
  clearTimeout(debounceTimer);
  debounceTimer = setTimeout(doGenerate, DEBOUNCE_MS);
}

async function doGenerate() {
  const form = document.getElementById('asciiForm');
  const pre = document.getElementById('asciiOutput');
  const err = document.getElementById('error');

  const text = form.inputText.value.trim();
  if (!text) return;

  err.hidden = true;
  err.innerHTML = '';
  pre.className = '';
  pre.innerHTML = '';

  const fd = new FormData();
  fd.append('inputText', text);
  fd.append('banner', form.banner.value);
  fd.append('align', form.align.value);
  fd.append('color', globalColorValue);

  const targets = form.colorTarget.value.split(',').map(s => s.trim()).filter(Boolean);
  targets.forEach(t => {
    fd.append('colorTarget', t);
    fd.append('targetColor', targetColorValue);
  });

  try {
    const res = await fetch('/ascii-art', { method: 'POST', body: fd });
    const html = await res.text();

    if (!res.ok) {
      err.innerHTML = html || '❌ Something went wrong.';
      err.hidden = false;
      return;
    }

    const wrapper = document.createElement('div');
    wrapper.className = 'inner';
    wrapper.innerHTML = html;
    pre.appendChild(wrapper);
    pre.classList.add('align-' + form.align.value);
  } catch (e) {
    err.textContent = 'Network error: ' + e.message;
    err.hidden = false;
  }
}
