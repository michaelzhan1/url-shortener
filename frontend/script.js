async function handleSubmit(e) {
  e.preventDefault();

  const formData = new FormData(e.target);

  const url = formData.get('url');
  const customId = formData.get('custom-id');

  let apiUrl;
  if (customId === "") {
    apiUrl = `http://localhost:8888/api/new?url=${encodeURIComponent(url)}`;
  } else {
    apiUrl = `http://localhost:8888/api/new/custom?url=${encodeURIComponent(url)}?id=${customId}`;
  }

  const response = await fetch(apiUrl);
  const id = await response.text();
  document.getElementById('shortened-url').innerHTML = `
    <div>Shortened URL</div>
    <div>http://localhost:8888/${id}</div>
  `;
}

document.getElementById('main-form').addEventListener('submit', handleSubmit);