require.config({ paths: { 'vs': 'https://unpkg.com/monaco-editor@0.16.2/min/vs' }});
window.MonacoEnvironment = { getWorkerUrl: () => proxy };

let proxy = URL.createObjectURL(new Blob([`
	self.MonacoEnvironment = {
		baseUrl: 'https://unpkg.com/monaco-editor@0.16.2/min/'
	};
	importScripts('https://unpkg.com/monaco-editor@0.16.2/min/vs/base/worker/workerMain.js');
`], { type: 'text/javascript' }));

let libContent=`
type noCodeSuggestionType =
    "noCodeSuggestion1" |
    "noCodeSuggestion2"

interface interfaceA {
    no_code_suggestion?: noCodeSuggestionType
}

interface interfaceB {
    interface_a?: interfaceA
}

interface TEST {
    find_by(attrs: interfaceB): void
    find_by_b(attrs: noCodeSuggestionType): void
}

declare var $test: TEST
`

require(["vs/editor/editor.main"], function () {
  let container = document.getElementById('container')
  let raw = container.dataset.raw
	let editor = monaco.editor.create(container, {
    value: raw,
		language: 'go',
    theme: 'vs-dark',
    fontSize: 20
  });
  
  window.editor = editor
	monaco.languages.typescript.typescriptDefaults.addExtraLib(libContent);
  monaco.languages.typescript.typescriptDefaults.setCompilerOptions({
    allowNonTsExtensions: true,
    noLib: true
  });
    
	// editor.addListener('didType', () => {
	// 	console.log(editor.getValue());
	// });
});