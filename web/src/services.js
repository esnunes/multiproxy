export async function loadSelectedEnv(addr) {
  const url = `${addr}/_multiproxy`;

  const res = await fetch(url, {
    method: 'GET',
    mode: 'cors',
    credentials: 'include',
  });

  if (!res.ok) {
    throw new Error(res.text() || res.statusText);
  }

  return res.json();
}

export async function selectEnv(addr, key) {
  const url = `${addr}/_multiproxy`;

  const res = await fetch(url, {
    method: 'POST',
    mode: 'cors',
    body: JSON.stringify({ key }),
    cache: 'no-cache',
    credentials: 'include',
    headers: {
      'content-type': 'application/json',
    },
  });

  if (!res.ok) {
    throw new Error(res.text() || res.statusText);
  }

  return res.json();
}
