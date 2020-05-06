package parser

import (
  "errors"
  "regexp"
  "strconv"

  "distributed-crawler-demo/config"
  "distributed-crawler-demo/engine"
  "distributed-crawler-demo/model"
)

var (
  ageRe           = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
  heightRe        = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
  weightRe        = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
  incomeRe        = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
  marriageRe      = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
  genderRe        = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
  constellationRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
  educationRe     = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
  occupationRe    = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
  residenceRe     = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
  houseRe         = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
  carRe           = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
  guessRe         = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(.*album\.zhenai\.com/u/[\d]+)">([^<]+)</a>`)
  idUrlRe         = regexp.MustCompile(`.*album\.zhenai\.com/u/([\d]+)`)
)

func parseProfile(contents []byte, url string, name string) engine.ParseResult {
  profile := model.Profile{}

  if age, err := extractInt(contents, ageRe); err == nil {
    profile.Age = age
  }
  if height, err := extractInt(contents, heightRe); err == nil {
    profile.Height = height
  }
  if weight, err := extractInt(contents, weightRe); err == nil {
    profile.Weight = weight
  }

  if income, err := extractString(contents, incomeRe); err == nil {
    profile.Income = income
  }
  if marriage, err := extractString(contents, marriageRe); err == nil {
    profile.Marriage = marriage
  }
  if gender, err := extractString(contents, genderRe); err == nil {
    profile.Gender = gender
  }
  if constellation, err := extractString(contents, constellationRe); err == nil {
    profile.Constellation = constellation
  }
  if education, err := extractString(contents, educationRe); err == nil {
    profile.Education = education
  }
  if occupation, err := extractString(contents, occupationRe); err == nil {
    profile.Occupation = occupation
  }
  if residence, err := extractString(contents, residenceRe); err == nil {
    profile.Residence = residence
  }
  if house, err := extractString(contents, houseRe); err == nil {
    profile.House = house
  }
  if car, err := extractString(contents, carRe); err == nil {
    profile.Car = car
  }

  id, _ := extractString([]byte(url), idUrlRe)

  profile.Name = name
  result := engine.ParseResult{
    Items: []engine.Item{
      {
        Url:     url,
        Id:      id,
        Payload: profile,
      },
    },
  }

  bytes := guessRe.FindAllSubmatch(contents, -1)
  for _, b := range bytes {
    result.Requests = append(result.Requests,
      engine.Request{
        Url:    string(b[1]),
        Parser: NewProfileParser(string(b[2])),
      })
  }
  return result
}

func extractString(contents []byte, re *regexp.Regexp) (string, error) {
  match := re.FindSubmatch(contents)
  if len(match) >= 2 {
    return string(match[1]), nil
  } else {
    return "", errors.New("extract string error")
  }
}

func extractInt(contents []byte, re *regexp.Regexp) (int, error) {
  digitString, err := extractString(contents, re)
  if err == nil {
    digit, err := strconv.Atoi(digitString)
    if err == nil {
      return digit, nil
    }
    return 0, err
  }
  return 0, err
}

type ProfileParser struct {
  userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
  return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
  return config.ParseProfile, p.userName
}

func NewProfileParser(name string) *ProfileParser {
  return &ProfileParser{
    userName: name,
  }
}
