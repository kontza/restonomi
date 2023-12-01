const fetchIt = (marker) => {
  console.log(`Sending request '${marker}'`);
  const retVal = fetch(`http://localhost:3210/api?id=${marker}`);
  console.log(`Promise for '${marker}'`);
  return retVal;
};
const withPromise = async () => {
  const values = await Promise.all([
    fetchIt("all 1"),
    fetchIt("all 2"),
    fetchIt("all 3"),
    fetchIt("all 4"),
  ]);
  values.forEach(async (v) => {
    console.log(await v.json());
  });
};

const withoutPromise = async () => {
  let r = await fetchIt("await 1");
  console.log(await r.json());
  r = await fetchIt("await 2");
  console.log(await r.json());
  r = await fetchIt("await 3");
  console.log(await r.json());
  r = await fetchIt("await 4");
  console.log(await r.json());
};
withPromise();
withoutPromise();
