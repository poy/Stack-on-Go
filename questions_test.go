package stackongo

import (
	"testing"
	"strings"
)

func TestAllQuestions(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/questions", dummyQuestionsResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	session := NewSession("stackoverflow")
	questions, err := session.AllQuestions(map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(questions) != 3 {
		t.Error("Number of items wrong.")
	}

	if questions[0].Title != "GetAvailableWebTemplates returns nothing" {
		t.Error("Title invalid.")
	}

	if strings.Join(questions[0].Tags, ", ") != "sharepoint, templates" {
		t.Error("Tags invalid.")
	}

	if questions[0].Owner.Display_name != "Nacht" {
		t.Error("Owner invalid.")
	}

	if questions[0].Creation_date != 1327453582 {
		t.Error("Date invalid.")
	}

}

func TestQuestions(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/questions/1;2;3", dummyQuestionsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.Questions([]int{1, 2, 3}, map[string]string{"sort": "votes", "order": "desc", "page": "1"})

	if err != nil {
		t.Error(err.String())
	}

}

func TestUnansweredQuestions(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/questions/unanswered", dummyQuestionsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.UnansweredQuestions(map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

}

func TestQuestionsWithNoAnswers(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/questions/no-answers", dummyQuestionsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.QuestionsWithNoAnswers(map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

}

func TestRelatedQuestions(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/questions/1;2;3/related", dummyQuestionsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.RelatedQuestions([]int{1, 2, 3}, map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

}

func TestLinkedQuestions(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/questions/1;2;3/linked", dummyQuestionsResponse, t)
	defer dummy_server.Close()

	session := NewSession("stackoverflow")
	_, err := session.LinkedQuestions([]int{1, 2, 3}, map[string]string{})

	if err != nil {
		t.Error(err.String())
	}

}

//Test Data

var dummyQuestionsResponse string = `{
  "items": [
    {
      "question_id": 8996647,
      "creation_date": 1327453582,
      "last_activity_date": 1327453582,
      "score": 0,
      "answer_count": 0,
      "title": "GetAvailableWebTemplates returns nothing",
      "tags": [
        "sharepoint",
        "templates"
      ],
      "view_count": 1,
      "owner": {
        "user_id": 983442,
        "display_name": "Nacht",
        "reputation": 132,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/5775b55f9b3be541dbbf9967cf3d1f16?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/983442/nacht"
      },
      "link": "http://stackoverflow.com/questions/8996647/getavailablewebtemplates-returns-nothing",
      "is_answered": false
    },
    {
      "question_id": 8996646,
      "creation_date": 1327453577,
      "last_activity_date": 1327453577,
      "score": 0,
      "answer_count": 0,
      "title": "Netbeans c++ remote execution",
      "tags": [
        "c",
        "netbeans",
        "remote-server"
      ],
      "view_count": 1,
      "owner": {
        "user_id": 676516,
        "display_name": "James Cotter",
        "reputation": 7,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/b3133db6c2e37607dd036766f3bcf6fd?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/676516/james-cotter"
      },
      "link": "http://stackoverflow.com/questions/8996646/netbeans-c-remote-execution",
      "is_answered": false
    },
    {
      "question_id": 8996578,
      "creation_date": 1327452958,
      "last_activity_date": 1327453572,
      "score": 0,
      "answer_count": 1,
      "title": "UrlEncode for Silverlight?",
      "tags": [
        "silverlight"
      ],
      "view_count": 4,
      "owner": {
        "user_id": 5274,
        "display_name": "Jonathan Allen",
        "reputation": 14239,
        "user_type": "registered",
        "profile_image": "http://www.gravatar.com/avatar/71c1027d2483fe242b0affc5e59df647?d=identicon&r=PG",
        "link": "http://stackoverflow.com/users/5274/jonathan-allen"
      },
      "link": "http://stackoverflow.com/questions/8996578/urlencode-for-silverlight",
      "is_answered": false
    }
  ],
  "quota_remaining": 9989,
  "quota_max": 10000,
  "has_more": true
}`
