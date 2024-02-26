package model

import "time"

type Quiz struct{
    ID  uint    `json:"id"`
    UserID  uint    `json:"user_id"`
    Title    string `json:"title"`
    StartTime   time.Time   `json:"start_time"`
    EndTime time.Time   `json:"end_time"`
    Questions   []Question  `json:"questions"`
    Results []Result    `json:"restults"`
}

type Question struct{
    ID  uint    `json:"question_id"`
    Text    string  `json:"text"`
    Options map[int]string  `json:"options"`
    Answer string   `json:"answer;omitempty"`
}

type Result struct{
    ID  uint    `json:"id"`
    QuizID  uint    `json:"quiz_id"`
    AllQuestions    uint    `json:"all"`
    CorrectAnswers    uint    `json:"correct;ommitempty"`
    Responses   []UserAnswer    `json:"responses"`
}

type UserAnswer struct {
    QuestionID   uint   `json:"question_id"`
    Answer       string `json:"answer"`
    IsCorrect    bool   `json:"is_correct,omitempty"`
}
