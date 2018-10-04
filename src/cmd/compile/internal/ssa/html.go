// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"bytes"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/src"
	"html"
	"io"
	"os"
	"strings"
)

type HTMLWriter struct {
	Logger
	w io.WriteCloser
}

func (pstate *PackageState) NewHTMLWriter(path string, logger Logger, funcname string) *HTMLWriter {
	out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		logger.Fatalf(pstate.src.NoXPos, "%v", err)
	}
	html := HTMLWriter{w: out, Logger: logger}
	html.start(pstate, funcname)
	return &html
}

func (w *HTMLWriter) start(pstate *PackageState, name string) {
	if w == nil {
		return
	}
	w.WriteString(pstate, "<html>")
	w.WriteString(pstate, "<head>\n<meta http-equiv=\"Content-Type\" content=\"text/html;charset=UTF-8\">\n<style>\n\nbody {\n    font-size: 14px;\n    font-family: Arial, sans-serif;\n}\n\n#helplink {\n    margin-bottom: 15px;\n    display: block;\n    margin-top: -15px;\n}\n\n#help {\n    display: none;\n}\n\n.stats {\n\tfont-size: 60%;\n}\n\ntable {\n    border: 1px solid black;\n    table-layout: fixed;\n    width: 300px;\n}\n\nth, td {\n    border: 1px solid black;\n    overflow: hidden;\n    width: 400px;\n    vertical-align: top;\n    padding: 5px;\n}\n\ntd > h2 {\n    cursor: pointer;\n    font-size: 120%;\n}\n\ntd.collapsed {\n    font-size: 12px;\n    width: 12px;\n    border: 0px;\n    padding: 0;\n    cursor: pointer;\n    background: #fafafa;\n}\n\ntd.collapsed  div {\n     -moz-transform: rotate(-90.0deg);  /* FF3.5+ */\n       -o-transform: rotate(-90.0deg);  /* Opera 10.5 */\n  -webkit-transform: rotate(-90.0deg);  /* Saf3.1+, Chrome */\n             filter:  progid:DXImageTransform.Microsoft.BasicImage(rotation=0.083);  /* IE6,IE7 */\n         -ms-filter: \"progid:DXImageTransform.Microsoft.BasicImage(rotation=0.083)\"; /* IE8 */\n         margin-top: 10.3em;\n         margin-left: -10em;\n         margin-right: -10em;\n         text-align: right;\n}\n\ntd.ssa-prog {\n    width: 600px;\n    word-wrap: break-word;\n}\n\nli {\n    list-style-type: none;\n}\n\nli.ssa-long-value {\n    text-indent: -2em;  /* indent wrapped lines */\n}\n\nli.ssa-value-list {\n    display: inline;\n}\n\nli.ssa-start-block {\n    padding: 0;\n    margin: 0;\n}\n\nli.ssa-end-block {\n    padding: 0;\n    margin: 0;\n}\n\nul.ssa-print-func {\n    padding-left: 0;\n}\n\ndl.ssa-gen {\n    padding-left: 0;\n}\n\ndt.ssa-prog-src {\n    padding: 0;\n    margin: 0;\n    float: left;\n    width: 4em;\n}\n\ndd.ssa-prog {\n    padding: 0;\n    margin-right: 0;\n    margin-left: 4em;\n}\n\n.dead-value {\n    color: gray;\n}\n\n.dead-block {\n    opacity: 0.5;\n}\n\n.depcycle {\n    font-style: italic;\n}\n\n.line-number {\n    font-style: italic;\n    font-size: 11px;\n}\n\n.highlight-aquamarine     { background-color: aquamarine; }\n.highlight-coral          { background-color: coral; }\n.highlight-lightpink      { background-color: lightpink; }\n.highlight-lightsteelblue { background-color: lightsteelblue; }\n.highlight-palegreen      { background-color: palegreen; }\n.highlight-skyblue        { background-color: skyblue; }\n.highlight-lightgray      { background-color: lightgray; }\n.highlight-yellow         { background-color: yellow; }\n.highlight-lime           { background-color: lime; }\n.highlight-khaki          { background-color: khaki; }\n.highlight-aqua           { background-color: aqua; }\n.highlight-salmon         { background-color: salmon; }\n\n.outline-blue           { outline: blue solid 2px; }\n.outline-red            { outline: red solid 2px; }\n.outline-blueviolet     { outline: blueviolet solid 2px; }\n.outline-darkolivegreen { outline: darkolivegreen solid 2px; }\n.outline-fuchsia        { outline: fuchsia solid 2px; }\n.outline-sienna         { outline: sienna solid 2px; }\n.outline-gold           { outline: gold solid 2px; }\n.outline-orangered      { outline: orangered solid 2px; }\n.outline-teal           { outline: teal solid 2px; }\n.outline-maroon         { outline: maroon solid 2px; }\n.outline-black          { outline: black solid 2px; }\n\n</style>\n\n<script type=\"text/javascript\">\n// ordered list of all available highlight colors\nvar highlights = [\n    \"highlight-aquamarine\",\n    \"highlight-coral\",\n    \"highlight-lightpink\",\n    \"highlight-lightsteelblue\",\n    \"highlight-palegreen\",\n    \"highlight-skyblue\",\n    \"highlight-lightgray\",\n    \"highlight-yellow\",\n    \"highlight-lime\",\n    \"highlight-khaki\",\n    \"highlight-aqua\",\n    \"highlight-salmon\"\n];\n\n// state: which value is highlighted this color?\nvar highlighted = {};\nfor (var i = 0; i < highlights.length; i++) {\n    highlighted[highlights[i]] = \"\";\n}\n\n// ordered list of all available outline colors\nvar outlines = [\n    \"outline-blue\",\n    \"outline-red\",\n    \"outline-blueviolet\",\n    \"outline-darkolivegreen\",\n    \"outline-fuchsia\",\n    \"outline-sienna\",\n    \"outline-gold\",\n    \"outline-orangered\",\n    \"outline-teal\",\n    \"outline-maroon\",\n    \"outline-black\"\n];\n\n// state: which value is outlined this color?\nvar outlined = {};\nfor (var i = 0; i < outlines.length; i++) {\n    outlined[outlines[i]] = \"\";\n}\n\nwindow.onload = function() {\n    var ssaElemClicked = function(elem, event, selections, selected) {\n        event.stopPropagation()\n\n        // TODO: pushState with updated state and read it on page load,\n        // so that state can survive across reloads\n\n        // find all values with the same name\n        var c = elem.classList.item(0);\n        var x = document.getElementsByClassName(c);\n\n        // if selected, remove selections from all of them\n        // otherwise, attempt to add\n\n        var remove = \"\";\n        for (var i = 0; i < selections.length; i++) {\n            var color = selections[i];\n            if (selected[color] == c) {\n                remove = color;\n                break;\n            }\n        }\n\n        if (remove != \"\") {\n            for (var i = 0; i < x.length; i++) {\n                x[i].classList.remove(remove);\n            }\n            selected[remove] = \"\";\n            return;\n        }\n\n        // we're adding a selection\n        // find first available color\n        var avail = \"\";\n        for (var i = 0; i < selections.length; i++) {\n            var color = selections[i];\n            if (selected[color] == \"\") {\n                avail = color;\n                break;\n            }\n        }\n        if (avail == \"\") {\n            alert(\"out of selection colors; go add more\");\n            return;\n        }\n\n        // set that as the selection\n        for (var i = 0; i < x.length; i++) {\n            x[i].classList.add(avail);\n        }\n        selected[avail] = c;\n    };\n\n    var ssaValueClicked = function(event) {\n        ssaElemClicked(this, event, highlights, highlighted);\n    }\n\n    var ssaBlockClicked = function(event) {\n        ssaElemClicked(this, event, outlines, outlined);\n    }\n\n    var ssavalues = document.getElementsByClassName(\"ssa-value\");\n    for (var i = 0; i < ssavalues.length; i++) {\n        ssavalues[i].addEventListener('click', ssaValueClicked);\n    }\n\n    var ssalongvalues = document.getElementsByClassName(\"ssa-long-value\");\n    for (var i = 0; i < ssalongvalues.length; i++) {\n        // don't attach listeners to li nodes, just the spans they contain\n        if (ssalongvalues[i].nodeName == \"SPAN\") {\n            ssalongvalues[i].addEventListener('click', ssaValueClicked);\n        }\n    }\n\n    var ssablocks = document.getElementsByClassName(\"ssa-block\");\n    for (var i = 0; i < ssablocks.length; i++) {\n        ssablocks[i].addEventListener('click', ssaBlockClicked);\n    }\n   var expandedDefault = [\n        \"start\",\n        \"deadcode\",\n        \"opt\",\n        \"lower\",\n        \"late deadcode\",\n        \"regalloc\",\n        \"genssa\",\n    ]\n    function isExpDefault(id) {\n        for (var i = 0; i < expandedDefault.length; i++) {\n            if (id.startsWith(expandedDefault[i])) {\n                return true;\n            }\n        }\n        return false;\n    }\n    function toggler(phase) {\n        return function() {\n            toggle_cell(phase+'-col');\n            toggle_cell(phase+'-exp');\n        };\n    }\n    function toggle_cell(id) {\n       var e = document.getElementById(id);\n       if(e.style.display == 'table-cell')\n          e.style.display = 'none';\n       else\n          e.style.display = 'table-cell';\n    }\n\n    var td = document.getElementsByTagName(\"td\");\n    for (var i = 0; i < td.length; i++) {\n        var id = td[i].id;\n        var def = isExpDefault(id);\n        var phase = id.substr(0, id.length-4);\n        if (id.endsWith(\"-exp\")) {\n            var h2 = td[i].getElementsByTagName(\"h2\");\n            if (h2 && h2[0]) {\n                h2[0].addEventListener('click', toggler(phase));\n            }\n        } else {\n\t        td[i].addEventListener('click', toggler(phase));\n        }\n        if (id.endsWith(\"-col\") && def || id.endsWith(\"-exp\") && !def) {\n               td[i].style.display = 'none';\n               continue\n        }\n        td[i].style.display = 'table-cell';\n    }\n};\n\nfunction toggle_visibility(id) {\n   var e = document.getElementById(id);\n   if(e.style.display == 'block')\n      e.style.display = 'none';\n   else\n      e.style.display = 'block';\n}\n</script>\n\n</head>")
	w.WriteString(pstate, "<body>")
	w.WriteString(pstate, "<h1>")
	w.WriteString(pstate, html.EscapeString(name))
	w.WriteString(pstate, "</h1>")
	w.WriteString(pstate, "\n<a href=\"#\" onclick=\"toggle_visibility('help');\" id=\"helplink\">help</a>\n<div id=\"help\">\n\n<p>\nClick on a value or block to toggle highlighting of that value/block\nand its uses.  (Values and blocks are highlighted by ID, and IDs of\ndead items may be reused, so not all highlights necessarily correspond\nto the clicked item.)\n</p>\n\n<p>\nFaded out values and blocks are dead code that has not been eliminated.\n</p>\n\n<p>\nValues printed in italics have a dependency cycle.\n</p>\n\n</div>\n")
	w.WriteString(pstate, "<table>")
	w.WriteString(pstate, "<tr>")
}

func (w *HTMLWriter) Close() {
	if w == nil {
		return
	}
	io.WriteString(w.w, "</tr>")
	io.WriteString(w.w, "</table>")
	io.WriteString(w.w, "</body>")
	io.WriteString(w.w, "</html>")
	w.w.Close()
}

// WriteFunc writes f in a column headed by title.
func (w *HTMLWriter) WriteFunc(pstate *PackageState, phase, title string, f *Func) {
	if w == nil {
		return // avoid generating HTML just to discard it
	}
	w.WriteColumn(pstate, phase, title, "", f.HTML(pstate))
	// TODO: Add visual representation of f's CFG.
}

// WriteColumn writes raw HTML in a column headed by title.
// It is intended for pre- and post-compilation log output.
func (w *HTMLWriter) WriteColumn(pstate *PackageState, phase, title, class, html string) {
	if w == nil {
		return
	}
	id := strings.Replace(phase, " ", "-", -1)
	// collapsed column
	w.Printf(pstate, "<td id=\"%v-col\" class=\"collapsed\"><div>%v</div></td>", id, phase)

	if class == "" {
		w.Printf(pstate, "<td id=\"%v-exp\">", id)
	} else {
		w.Printf(pstate, "<td id=\"%v-exp\" class=\"%v\">", id, class)
	}
	w.WriteString(pstate, "<h2>"+title+"</h2>")
	w.WriteString(pstate, html)
	w.WriteString(pstate, "</td>")
}

func (w *HTMLWriter) Printf(pstate *PackageState, msg string, v ...interface{}) {
	if _, err := fmt.Fprintf(w.w, msg, v...); err != nil {
		w.Fatalf(pstate.src.NoXPos, "%v", err)
	}
}

func (w *HTMLWriter) WriteString(pstate *PackageState, s string) {
	if _, err := io.WriteString(w.w, s); err != nil {
		w.Fatalf(pstate.src.NoXPos, "%v", err)
	}
}

func (v *Value) HTML() string {
	// TODO: Using the value ID as the class ignores the fact
	// that value IDs get recycled and that some values
	// are transmuted into other values.
	s := v.String()
	return fmt.Sprintf("<span class=\"%s ssa-value\">%s</span>", s, s)
}

func (v *Value) LongHTML(pstate *PackageState) string {
	// TODO: Any intra-value formatting?
	// I'm wary of adding too much visual noise,
	// but a little bit might be valuable.
	// We already have visual noise in the form of punctuation
	// maybe we could replace some of that with formatting.
	s := fmt.Sprintf("<span class=\"%s ssa-long-value\">", v.String())

	linenumber := "<span class=\"line-number\">(?)</span>"
	if v.Pos.IsKnown() {
		linenumber = fmt.Sprintf("<span class=\"line-number\">(%s)</span>", v.Pos.LineNumberHTML())
	}

	s += fmt.Sprintf("%s %s = %s", v.HTML(), linenumber, v.Op.String(pstate))

	s += " &lt;" + html.EscapeString(v.Type.String(pstate.types)) + "&gt;"
	s += html.EscapeString(v.auxString(pstate))
	for _, a := range v.Args {
		s += fmt.Sprintf(" %s", a.HTML())
	}
	r := v.Block.Func.RegAlloc
	if int(v.ID) < len(r) && r[v.ID] != nil {
		s += " : " + html.EscapeString(r[v.ID].String())
	}
	var names []string
	for name, values := range v.Block.Func.NamedValues {
		for _, value := range values {
			if value == v {
				names = append(names, name.String())
				break // drop duplicates.
			}
		}
	}
	if len(names) != 0 {
		s += " (" + strings.Join(names, ", ") + ")"
	}

	s += "</span>"
	return s
}

func (b *Block) HTML() string {
	// TODO: Using the value ID as the class ignores the fact
	// that value IDs get recycled and that some values
	// are transmuted into other values.
	s := html.EscapeString(b.String())
	return fmt.Sprintf("<span class=\"%s ssa-block\">%s</span>", s, s)
}

func (b *Block) LongHTML(pstate *PackageState) string {
	// TODO: improve this for HTML?
	s := fmt.Sprintf("<span class=\"%s ssa-block\">%s</span>", html.EscapeString(b.String()), html.EscapeString(b.Kind.String(pstate)))
	if b.Aux != nil {
		s += html.EscapeString(fmt.Sprintf(" {%v}", b.Aux))
	}
	if b.Control != nil {
		s += fmt.Sprintf(" %s", b.Control.HTML())
	}
	if len(b.Succs) > 0 {
		s += " &#8594;" // right arrow
		for _, e := range b.Succs {
			c := e.b
			s += " " + c.HTML()
		}
	}
	switch b.Likely {
	case BranchUnlikely:
		s += " (unlikely)"
	case BranchLikely:
		s += " (likely)"
	}
	if b.Pos.IsKnown() {
		// TODO does not begin to deal with the full complexity of line numbers.
		// Maybe we want a string/slice instead, of outer-inner when inlining.
		s += fmt.Sprintf(" (line %s)", b.Pos.LineNumberHTML())
	}
	return s
}

func (f *Func) HTML(pstate *PackageState) string {
	var buf bytes.Buffer
	fmt.Fprint(&buf, "<code>")
	p := htmlFuncPrinter{w: &buf}
	pstate.fprintFunc(p, f)

	// fprintFunc(&buf, f) // TODO: HTML, not text, <br /> for line breaks, etc.
	fmt.Fprint(&buf, "</code>")
	return buf.String()
}

type htmlFuncPrinter struct {
	w io.Writer
}

func (p htmlFuncPrinter) header(f *Func) {}

func (p htmlFuncPrinter) startBlock(b *Block, reachable bool) {
	// TODO: Make blocks collapsable?
	var dead string
	if !reachable {
		dead = "dead-block"
	}
	fmt.Fprintf(p.w, "<ul class=\"%s ssa-print-func %s\">", b, dead)
	fmt.Fprintf(p.w, "<li class=\"ssa-start-block\">%s:", b.HTML())
	if len(b.Preds) > 0 {
		io.WriteString(p.w, " &#8592;") // left arrow
		for _, e := range b.Preds {
			pred := e.b
			fmt.Fprintf(p.w, " %s", pred.HTML())
		}
	}
	io.WriteString(p.w, "</li>")
	if len(b.Values) > 0 { // start list of values
		io.WriteString(p.w, "<li class=\"ssa-value-list\">")
		io.WriteString(p.w, "<ul>")
	}
}

func (p htmlFuncPrinter) endBlock(pstate *PackageState, b *Block) {
	if len(b.Values) > 0 { // end list of values
		io.WriteString(p.w, "</ul>")
		io.WriteString(p.w, "</li>")
	}
	io.WriteString(p.w, "<li class=\"ssa-end-block\">")
	fmt.Fprint(p.w, b.LongHTML(pstate))
	io.WriteString(p.w, "</li>")
	io.WriteString(p.w, "</ul>")
	// io.WriteString(p.w, "</span>")
}

func (p htmlFuncPrinter) value(pstate *PackageState, v *Value, live bool) {
	var dead string
	if !live {
		dead = "dead-value"
	}
	fmt.Fprintf(p.w, "<li class=\"ssa-long-value %s\">", dead)
	fmt.Fprint(p.w, v.LongHTML(pstate))
	io.WriteString(p.w, "</li>")
}

func (p htmlFuncPrinter) startDepCycle() {
	fmt.Fprintln(p.w, "<span class=\"depcycle\">")
}

func (p htmlFuncPrinter) endDepCycle() {
	fmt.Fprintln(p.w, "</span>")
}

func (p htmlFuncPrinter) named(n LocalSlot, vals []*Value) {
	fmt.Fprintf(p.w, "<li>name %s: ", n)
	for _, val := range vals {
		fmt.Fprintf(p.w, "%s ", val.HTML())
	}
	fmt.Fprintf(p.w, "</li>")
}
