package utils

import (
  "net/http"
  "strconv"
	"github.com/gin-gonic/gin"
)

const TEMP_FILE = "tmp/tmp_data.csv"

type EventsList struct {
  previous_point []int
  next_point []int
  transition_status []int
}

// Returns list of unique values
func unique(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{}
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}


// Returns successes and fails transition matricies
func (el *EventsList) BuildTransitionMatrix() ([][]float32, [][]float32) {
  from_points := unique(el.previous_point)
  to_points := unique(el.next_point)
  n_points_unique := len(unique(append(from_points, to_points...)))

  n_events := len(el.transition_status)

  // build square matrix for all points and fill values by 0
  transition_matrix_success := make([][]float32, n_points_unique)
  transition_matrix_fail := make([][]float32, n_points_unique)
  for i := 0; i < n_points_unique; i++ {
    transition_matrix_success[i] = make([]float32, n_points_unique)
    transition_matrix_fail[i] = make([]float32, n_points_unique)
  }

  // count occurences for points transitions
  for i := 0; i < n_events; i++ {
    fp := el.previous_point[i] - 1
    np := el.next_point[i] - 1
    ts := el.transition_status[i] - 1

    if ts == 0 {
      transition_matrix_fail[fp][np] += 1.0 / float32(n_points_unique)
    } else {
      transition_matrix_success[fp][np] += 1.0 / float32(n_points_unique)
    }
  }

  return transition_matrix_success, transition_matrix_fail
}

func OptrainEventsFromCsv(c *gin.Context) {
    var previous_point []int
    var next_point []int
    var transition_status []int

    // write file
    file, _ := c.FormFile("file")
    c.SaveUploadedFile(file, TEMP_FILE)

    // read data
    records := ReadCsvFile(TEMP_FILE)

    // group data
    for index, values := range records {
      if index == 0 {
        continue
      }
      pp, _ := strconv.Atoi(values[0])
      np, _ := strconv.Atoi(values[1])
      tp, _ := strconv.Atoi(values[2])

      previous_point = append(previous_point, pp)
      next_point = append(next_point, np)
      transition_status = append(transition_status, tp)
    }

    // create frame
    events := EventsList{
      previous_point: previous_point,
      next_point: next_point,
      transition_status: transition_status,
    }

    // build matricies
    succMat, failMat := events.BuildTransitionMatrix()

    // return answer
    answer := map[string][][]float32{
      "success_matrix" : succMat,
      "fail_matrix": failMat,
    }
    c.JSON(http.StatusOK, answer)
}
