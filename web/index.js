const go = new Go();
WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject).then((result) => {
  go.run(result.instance);
  const password = pwgen();
  console.log(password);
});
