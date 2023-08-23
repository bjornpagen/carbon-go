package carbon

import (
	"bytes"
	"testing"
)

func TestAccordion(t *testing.T) {
	expected := `<ul class="bx--accordion bx--accordion--end"></ul>`

	b := &bytes.Buffer{}
	Accordion().Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestAccordionWithItems(t *testing.T) {
	expected := `<ul class="bx--accordion bx--accordion--end"><li class="bx--accordion__item"><button type="button" class="bx--accordion__heading" title="Expand/Collapse" aria-expanded="false"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" role="img" class="bx--accordion__arrow" aria-label="Expand/Collapse"><path d="M22 16L12 26 10.6 24.6 19.2 16 10.6 7.4 12 6z"></path></svg><div class="bx--accordion__title">Natural Language Classifier</div></button><div class="bx--accordion__content"><p>Natural Language Classifier uses advanced natural language processing and machine learning techniques to create custom classification models. Users train their data and the service predicts the appropriate category for the inputted text.</p></div></li><li class="bx--accordion__item"><button type="button" class="bx--accordion__heading" title="Expand/Collapse" aria-expanded="false"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" role="img" class="bx--accordion__arrow" aria-label="Expand/Collapse"><path d="M22 16L12 26 10.6 24.6 19.2 16 10.6 7.4 12 6z"></path></svg><div class="bx--accordion__title">Natural Language Understanding</div></button><div class="bx--accordion__content"><p>Analyze text to extract meta-data from content such as concepts, entities, emotion, relations, sentiment and more.</p></div></li><li class="bx--accordion__item"><button type="button" class="bx--accordion__heading" title="Expand/Collapse" aria-expanded="false"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" role="img" class="bx--accordion__arrow" aria-label="Expand/Collapse"><path d="M22 16L12 26 10.6 24.6 19.2 16 10.6 7.4 12 6z"></path></svg><div class="bx--accordion__title">Language Translator</div></button><div class="bx--accordion__content"><p>Translate text, documents, and websites from one language to another. Create industry or region-specific translations via the service's customization capability.</p></div></li></ul>`

	b := &bytes.Buffer{}
	Accordion(
		AccordionItem().Title("Natural Language Classifier").Content(
			Raw(`<p>Natural Language Classifier uses advanced natural language processing and machine learning techniques to create custom classification models. Users train their data and the service predicts the appropriate category for the inputted text.</p>`),
		),
		AccordionItem().Title("Natural Language Understanding").Content(
			Raw(`<p>Analyze text to extract meta-data from content such as concepts, entities, emotion, relations, sentiment and more.</p>`),
		),
		AccordionItem().Title("Language Translator").Content(
			Raw(`<p>Translate text, documents, and websites from one language to another. Create industry or region-specific translations via the service's customization capability.</p>`),
		),
	).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestAccordionCustomTitle(t *testing.T) {
	expected := `<ul class="bx--accordion bx--accordion--end"><li class="bx--accordion__item"><button type="button" class="bx--accordion__heading" title="Expand/Collapse" aria-expanded="false"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" role="img" class="bx--accordion__arrow" aria-label="Expand/Collapse"><path d="M22 16L12 26 10.6 24.6 19.2 16 10.6 7.4 12 6z"></path></svg><div class="bx--accordion__title"><h5>Natural Language Classifier</h5><p>AI / Machine Learning</p></div></button><div class="bx--accordion__content"><p>Natural Language Classifier uses advanced natural language processing and machine learning techniques to create custom classification models. Users train their data and the service predicts the appropriate category for the inputted text.</p></div></li><li class="bx--accordion__item"><button type="button" class="bx--accordion__heading" title="Expand/Collapse" aria-expanded="false"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" role="img" class="bx--accordion__arrow" aria-label="Expand/Collapse"><path d="M22 16L12 26 10.6 24.6 19.2 16 10.6 7.4 12 6z"></path></svg><div class="bx--accordion__title"><h5>Natural Language Understanding</h5><p>AI / Machine Learning</p></div></button><div class="bx--accordion__content"><p>Analyze text to extract meta-data from content such as concepts, entities, emotion, relations, sentiment and more.</p></div></li><li class="bx--accordion__item"><button type="button" class="bx--accordion__heading" title="Expand/Collapse" aria-expanded="false"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" role="img" class="bx--accordion__arrow" aria-label="Expand/Collapse"><path d="M22 16L12 26 10.6 24.6 19.2 16 10.6 7.4 12 6z"></path></svg><div class="bx--accordion__title"><h5>Language Translator</h5><p>AI / Machine Learning</p></div></button><div class="bx--accordion__content"><p>Translate text, documents, and websites from one language to another. Create industry or region-specific translations via the service's customization capability.</p></div></li></ul>`

	b := &bytes.Buffer{}
	Accordion(
		AccordionItem().TitleComponent(
			H5(Raw("Natural Language Classifier")),
			P(Raw("AI / Machine Learning")),
		).Content(
			P(Raw("Natural Language Classifier uses advanced natural language processing and machine learning techniques to create custom classification models. Users train their data and the service predicts the appropriate category for the inputted text.")),
		),
		AccordionItem().TitleComponent(
			H5(Raw("Natural Language Understanding")),
			P(Raw("AI / Machine Learning")),
		).Content(
			P(Raw("Analyze text to extract meta-data from content such as concepts, entities, emotion, relations, sentiment and more.")),
		),
		AccordionItem().TitleComponent(
			H5(Raw("Language Translator")),
			P(Raw("AI / Machine Learning")),
		).Content(
			P(Raw("Translate text, documents, and websites from one language to another. Create industry or region-specific translations via the service's customization capability.")),
		),
	).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func TestAccordionFirstItemOpen(t *testing.T) {
	expected := `<ul class="bx--accordion bx--accordion--end"><li class="bx--accordion__item bx--accordion__item--active"><button type="button" class="bx--accordion__heading" title="Expand/Collapse" aria-expanded="true"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" role="img" class="bx--accordion__arrow" aria-label="Expand/Collapse"><path d="M22 16L12 26 10.6 24.6 19.2 16 10.6 7.4 12 6z"></path></svg><div class="bx--accordion__title">Natural Language Classifier</div></button><div class="bx--accordion__content"><p>Natural Language Classifier uses advanced natural language processing and machine learning techniques to create custom classification models. Users train their data and the service predicts the appropriate category for the inputted text.</p></div></li><li class="bx--accordion__item"><button type="button" class="bx--accordion__heading" title="Expand/Collapse" aria-expanded="false"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" role="img" class="bx--accordion__arrow" aria-label="Expand/Collapse"><path d="M22 16L12 26 10.6 24.6 19.2 16 10.6 7.4 12 6z"></path></svg><div class="bx--accordion__title">Natural Language Understanding</div></button><div class="bx--accordion__content"><p>Analyze text to extract meta-data from content such as concepts, entities, emotion, relations, sentiment and more.</p></div></li><li class="bx--accordion__item"><button type="button" class="bx--accordion__heading" title="Expand/Collapse" aria-expanded="false"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill="currentColor" preserveAspectRatio="xMidYMid meet" width="16" height="16" role="img" class="bx--accordion__arrow" aria-label="Expand/Collapse"><path d="M22 16L12 26 10.6 24.6 19.2 16 10.6 7.4 12 6z"></path></svg><div class="bx--accordion__title">Language Translator</div></button><div class="bx--accordion__content"><p>Translate text, documents, and websites from one language to another. Create industry or region-specific translations via the service's customization capability.</p></div></li></ul>`

	b := &bytes.Buffer{}
	Accordion(
		AccordionItem().Title("Natural Language Classifier").Content(
			Raw(`<p>Natural Language Classifier uses advanced natural language processing and machine learning techniques to create custom classification models. Users train their data and the service predicts the appropriate category for the inputted text.</p>`),
		).Open(true),
		AccordionItem().Title("Natural Language Understanding").Content(
			Raw(`<p>Analyze text to extract meta-data from content such as concepts, entities, emotion, relations, sentiment and more.</p>`),
		),
		AccordionItem().Title("Language Translator").Content(
			Raw(`<p>Translate text, documents, and websites from one language to another. Create industry or region-specific translations via the service's customization capability.</p>`),
		),
	).Render(b)
	if b.String() != expected {
		t.Errorf("got %s, want %s", b.String(), expected)
	}
}

func BenchmarkAccordionCustomTitleInitialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Accordion(
			AccordionItem().TitleComponent(
				H5(Raw("Natural Language Classifier")),
				P(Raw("AI / Machine Learning")),
			).Content(
				P(Raw("Natural Language Classifier uses advanced natural language processing and machine learning techniques to create custom classification models. Users train their data and the service predicts the appropriate category for the inputted text.")),
			),
			AccordionItem().TitleComponent(
				H5(Raw("Natural Language Understanding")),
				P(Raw("AI / Machine Learning")),
			).Content(
				P(Raw("Analyze text to extract meta-data from content such as concepts, entities, emotion, relations, sentiment and more.")),
			),
			AccordionItem().TitleComponent(
				H5(Raw("Language Translator")),
				P(Raw("AI / Machine Learning")),
			).Content(
				P(Raw("Translate text, documents, and websites from one language to another. Create industry or region-specific translations via the service's customization capability.")),
			),
		)
	}
}

func BenchmarkAccordionCustomTitle(b *testing.B) {
	a := Accordion(
		AccordionItem().TitleComponent(
			H5(Raw("Natural Language Classifier")),
			P(Raw("AI / Machine Learning")),
		).Content(
			P(Raw("Natural Language Classifier uses advanced natural language processing and machine learning techniques to create custom classification models. Users train their data and the service predicts the appropriate category for the inputted text.")),
		),
		AccordionItem().TitleComponent(
			H5(Raw("Natural Language Understanding")),
			P(Raw("AI / Machine Learning")),
		).Content(
			P(Raw("Analyze text to extract meta-data from content such as concepts, entities, emotion, relations, sentiment and more.")),
		),
		AccordionItem().TitleComponent(
			H5(Raw("Language Translator")),
			P(Raw("AI / Machine Learning")),
		).Content(
			P(Raw("Translate text, documents, and websites from one language to another. Create industry or region-specific translations via the service's customization capability.")),
		),
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		a.Render(&bytes.Buffer{})
	}
}
