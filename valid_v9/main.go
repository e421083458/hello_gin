package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_tw_translations "gopkg.in/go-playground/validator.v9/translations/zh_tw"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

type User struct {
	Username string `form:"user_name" validate:"required"`
	Tagline  string `form:"tag_line" validate:"required,lt=10"`
	Tagline2 string `form:"tag_line2" validate:"required,gt=1"`
}

func main() {
	en := en.New()
	zh := zh.New()
	zh_tw := zh_Hant_TW.New()
	Uni = ut.New(en, zh, zh_tw)
	Validate = validator.New()

	route := gin.Default()
	route.GET("/testing", startPage)
	route.POST("/testing", startPage)
	route.Run(":8080")
}

func startPage(c *gin.Context) {
	//这部分应放到中间件中
	locale:=c.DefaultQuery("locale","zh")
	trans, _ := Uni.GetTranslator(locale)
	switch locale {
	case "zh":
		zh_translations.RegisterDefaultTranslations(Validate, trans)
		break
	case "en":
		en_translations.RegisterDefaultTranslations(Validate, trans)
		break
	case "zh_tw":
		zh_tw_translations.RegisterDefaultTranslations(Validate, trans)
		break
	default:
		zh_translations.RegisterDefaultTranslations(Validate, trans)
		break
	}

	//自定义错误内容
	Validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	//这块应该放到公共验证方法中
	user := User{}
	c.ShouldBind(&user)
	fmt.Println(user)
	err := Validate.Struct(user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs:=[]string{}
		for _, e := range errs {
			sliceErrs=append(sliceErrs,e.Translate(trans))
		}
		c.String(200,fmt.Sprintf("%#v",sliceErrs))
	}
	c.String(200, fmt.Sprintf("%#v",""))
}