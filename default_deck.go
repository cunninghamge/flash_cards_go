package main

var card1 = Card{
	Question: "What is the capital of Alaska?",
	Answer:   "Juneau",
	Category: "Geography",
}
var card2 = Card{
	Question: "The Viking spacecraft sent back to Earth photographs and reports about the surface of which planet?",
	Answer:   "Mars",
	Category: "STEM",
}
var card3 = Card{
	Question: "Describe in words the exact direction that is 697.5° clockwise from due north?",
	Answer:   "North north west",
	Category: "STEM",
}
var card4 = Card{
	Question: "Which NFL position has, on average, the shortest career?",
	Answer:   "running back",
	Category: "Sports",
}

var card5 = Card{
	Question: "Where were the Utah Jazz originally located?",
	Answer:   "new orleans",
	Category: "Sports",
}

var card6 = Card{
	Question: "What year was the World Wide Web launched in?",
	Answer:   "1991",
	Category: "STEM",
}

var card7 = Card{
	Question: "What animal has the highest blood pressure?",
	Answer:   "giraffe",
	Category: "STEM",
}

var card8 = Card{
	Question: "Which Queen of England and wife of Henry VIII had six fingers on one of her hands?",
	Answer:   "anne boleyn",
	Category: "History",
}

var card9 = Card{
	Question: "What was the former name of Times Square before it was renamed in 1904?",
	Answer:   "longacre square",
	Category: "History",
}

var defaultDeck = Deck{[]Card{card1, card2, card3, card4, card5, card6, card7, card8, card9}}
