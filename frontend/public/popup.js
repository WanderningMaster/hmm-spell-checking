document.getElementById('openTab').addEventListener('click', function() {
  chrome.tabs.create({url: 'index.html'});
});
