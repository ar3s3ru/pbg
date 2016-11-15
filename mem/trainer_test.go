package mem

import (
    "testing"
    "encoding/json"
    "bytes"
)

type trainerCaseTest struct {
    trainer *Trainer
    result  []byte
}

var (
    trainer1 = &Trainer{}
)

func TestTrainer_MarshalJSON(t *testing.T) {
    fakeTr1, err := NewTrainer(WithTrainerName("test"), WithTrainerPassword("test"))
    if err != nil {
        t.Fatalf("Error occurred while creating: %s\n", err.Error())
    }

    var (
        fakeTrainer1 = fakeTr1.(*Trainer)

        trainerCaseTests = []trainerCaseTest{
            {
                trainer1,
                []byte(`{"name":"test","sign_up":"` + fakeTrainer1.SignUpDate().String() + `","set":false}`),
            },
            //{&trainer{}, []byte(``)},
        }
    )

    for _, test := range trainerCaseTests {
        m, err := json.Marshal(test.trainer)
        if err != nil {
            t.Fatalf("Error occurred while marshaling: %s\n", err.Error())
        }

        if !bytes.Equal(m, test.result) {
            t.Errorf("Result and test expectation are different:\n\tGot:%s\n\tExpect:%s\n", m, test.result)
        }
    }
}
