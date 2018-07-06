package main

import (
	"regexp"
	"sort"
)

// Table source: https://www.cl.cam.ac.uk/~mgk25/ucs/examples/TeX.txt
// slightly modified.

var gen = `
Lowercase Greek letters
	α \alpha	ι \iota		ϱ \varrho
	β \beta		κ \kappa	σ \sigma
	γ \gamma	λ \lambda	ς \varsigma
	δ \delta	μ \mu		τ \tau
	ϵ \epsilon	ν \nu		υ \upsilon
	ε \varepsilon	ξ \xi		ϕ \phi
	ζ \zeta	ο \o	φ \varphi
	η \eta	π \pi	χ \chi
	θ \theta	ϖ \varpi	ψ \psi
	ϑ \vartheta	ρ \rho		ω \omega

Uppercase Greek letters
	Γ \Gamma	Ξ \Xi		Φ \Phi
	Δ \Delta	Π \Pi		Ψ \Psi
	Θ \Theta	Σ \Sigma	Ω \Omega
	Λ \Lambda	Υ \Upsilon

Miscellaneous symbols
	ℵ \aleph	′ \prime	∀ \forall
	ℏ \hbar		∅ \emptyset	∃ \exists
	ı \imath	∇ \nabla	¬ \neg		¬ \lnot
	j \jmath	√ \surd		♭ \flat
	ℓ \ell		⊤ \top		♮ \natural
	℘ \wp		⊥ \bot		♯ \sharp
	ℜ \Re		∥ \|		♣ \clubsuit	∥ \Vert
	ℑ \Im		∠ \angle	♢ \diamondsuit
	∂ \partial	△ \triangle	♡ \heartsuit
	∞ \infty	\ \backslash	♠ \spadesuit
	□ \Box		◇ \Diamond

Large operators
	∑ \sum		⋂ \bigcap	⨀ \bigodot
	∏ \prod		⋃ \bigcup	⨂ \bigotimes
	∐ \coprod	⨆ \bigsqcup	⨁ \bigoplus
	∫ \int		⋁ \bigvee	⨄ \biguplus
	∮ \oint		⋀ \bigwedge

Binary operations
	± \pm		∩ \cap			∨ \vee		∨ \lor
	∓ \mp		∪ \cup			∧ \wedge	∧ \land
	∖ \setminus	⊎ \uplus		⊕ \oplus
	⋅ \cdot		⊓ \sqcap		⊖ \ominus
	× \times	⊔ \sqcup		⊗ \otimes
	∗ \ast		◁ \triangleleft		⊘ \oslash
	⋆ \star		▷ \triangleright	⊙ \odot
	⋄ \diamond	≀ \wr			† \dagger
	∘ \circ		◯ \bigcirc		‡ \ddagger
	∙ \bullet	△ \bigtriangleup	⨿ \amalg
	÷ \div		▽ \bigtriangledown	⊴ \unlhd
	⊲ \lhd		⊳ \rhd			⊵ \unrhd

Relations
	≤ \leq		≥ \geq		≡ \equiv	≤ \le	≥ \ge
	≺ \prec		≻ \succ		∼ \sim
	≼ \preceq	≽ \succeq	≃ \simeq
	≪ \ll		≫ \gg		≍ \asymp
	⊂ \subset	⊃ \supset	≈ \approx
	⊆ \subseteq	⊇ \supseteq	≅ \cong
	⊑ \sqsubseteq	⊒ \sqsupseteq	⋈ \bowtie
	∈ \in		∋ \ni		∝ \propto	∋ \owns
	⊢ \vdash	⊣ \dashv	⊨ \models
	⌣ \smile	∣ \mid		≐ \doteq	∣ \vert
	⌢ \frown	∥ \parallel	⊥ \perp
	⊏ \sqsubset	⊐ \sqsupset	⨝ \Join

Negated relations
	≮ \not<			≯ \not>			≠ \not=		≠ \ne
	≰ \notleq		≱ \notgeq		≢ \notequiv	≠ \neq
	⊀ \notprec		⊁ \notsucc		≁ \notsim
	⋠ \notpreceq		⋡ \notsucceq		≄ \notsimeq
	⊄ \notsubset		⊅ \notsupset		≉ \notapprox
	⊈ \notsubseteq		⊉ \notsupseteq		≇ \notcong
	⋢ \notsqsubseteq 	⋣ \notsqsupseteq	≭ \notasymp

Arrows
	← \leftarrow		⟵ \longleftarrow	↑ \uparrow	→ \to
	⇐ \Leftarrow		⟸ \Longleftarrow	⇑ \Uparrow	← \gets
	→ \rightarrow		⟶ \longrightarrow	↓ \downarrow
	⇒ \Rightarrow		⟹ \Longrightarrow	⇓ \Downarrow
	↔ \leftrightarrow	⟷ \longleftrightarrow	↕ \updownarrow
	⇔ \Leftrightarrow	⟺ \Longleftrightarrow	⇕ \Updownarrow
	↦ \mapsto		⟼ \longmapsto		↗ \nearrow
	↩ \hookleftarrow	↪ \hookrightarrow	↘ \searrow
	↼ \leftharpoonup	⇀ \rightharpoonup	↙ \swarrow
	↽ \leftharpoondown	⇁ \rightharpoondown	↖ \nwarrow
	⇌ \rightleftharpoons	↝ \leadsto

Openings/Closings
	[ \lbrack	⌊ \lfloor	⌈ \lceil	⟦ \[[
	{ \lbrace	⟨ \langle	⟪ \llangle
	] \rbrack	⌋ \rfloor	⌉ \rceil	⟧ \]]
	} \rbrace	⟩ \rangle	⟫ \rrangle
Dots
	⋮ \vdots
	⋯ \cdots
	⋱ \ddots
`

// Table stores a list of unicode replacement in [0], and patterns in [1].
var table [][2]string

func init() {
	// Generate the replacement table from gen.
	// We recognize a new replacement pattern following at least one tab.
	// replacement and pattern are separated by a blank and the pattern ends
	// with a tab or a newline.
	reg := regexp.MustCompile(`\t\t*([^ ]+) \\([^\t\n]+)`)
	for _, e := range reg.FindAllStringSubmatch(gen, -1) {
		table = append(table, [2]string{e[1], `\` + e[2]})
	}

	// We sort the table backwards, such that \alpha matches before \al.
	sort.Slice(table, func(i, j int) bool { return table[i][1] > table[j][1] })
}
