require.config({ paths: { 'vs': 'https://unpkg.com/monaco-editor@0.8.3/min/vs' }});
window.MonacoEnvironment = { getWorkerUrl: () => proxy };

let proxy = URL.createObjectURL(new Blob([`
	self.MonacoEnvironment = {
		baseUrl: 'https://unpkg.com/monaco-editor@0.8.3/min/'
	};
	importScripts('https://unpkg.com/monaco-editor@0.8.3/min/vs/base/worker/workerMain.js');
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
	let editor = monaco.editor.create(document.getElementById('container'), {
		value: [`$test.find_by({interface_a: {no_code_suggestion: ""}})`, `$test.find_by_b("")`].join('\n'),
		language: 'typescript',
		theme: 'vs-dark'
  });
  console.log(editor)
	monaco.languages.typescript.typescriptDefaults.addExtraLib(libContent);
  monaco.languages.typescript.typescriptDefaults.setCompilerOptions({
    allowNonTsExtensions: true,
    noLib: true
  });
    
	editor.addListener('didType', () => {
		console.log(editor.getValue());
	});
});