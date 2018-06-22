package main

import (
	"dpm/common"
	"time"
)

const (
	//简历仅自己可见
	CV_VISIBILI_TYPE_PRIVATE = 0
	//简历对所有人可见
	CV_VISIBILI_TYPE_PUBLIC = 1
	//简历用密码可见
	CV_VISIBILI_TYPE_PWD = 2

	//基本信息模块类型
	CVMTYPE_BASICINFO = "cvmtype_basicinfo"
	//头像模块类型
	CVMTYPE_HEADPORTRAIT = "cvmtype_headportrait"
	//求职意向模块类型
	CVMTYPE_JOBINTENSION = "cvmtype_job_intension"
	//项目经验模块类型
	CVMTYPE_PROJECTEXPERIENCE = "cvmtype_project_experience"
	//教育经历模块类型
	CVMTYPE_EDUCATEEXPERIENCE = "cvmtype_educate_experience"
	//自我评价模块类型
	CVMTYPE_SELFEVALUATION = "cvmtype_self_evaluation"
)

//简历
type CurriculumVitae struct {
	//简历Id
	CVId string
	//简历名称
	CVName string
	//查看密码
	CViewPwd string
	//自定义域名
	CustomDomainName string
	//可见类型
	CVisibiliType int
	//创建时间
	CVCreateTime common.JSONTime
	//更新时间
	CVUpdateTime common.JSONTime
	//简历模块
	CVModules map[string]interface{}
}

//简历模块
type CVModule struct {
	//模块类型
	CVMType string
	//模块名称
	CVMName string
	//是否显示
	ModuleDisplay bool
	//是否必须
	ModuleRequisite bool
	//坐标位置
	ModuleCoordinate string
}

//模块标题
type TitleCVM struct {
	*CVModule
	//模块标题
	TitleName string
	//标题显示
	TitleDisplay bool
}

//基本信息模块
type BasicInfoCVM struct {
	*TitleCVM
	//候选人姓名
	CandidateName string
	//出生年份
	YearOfBirth int
	//出生月份
	MonthOfBirth int
	//所在城市
	CityLocation string
	//工作年限
	WorkingLife int
	//联系电话
	PhoneNumber string
	//邮箱地址
	EMailAddress string
	//简短描述
	ShortDescribe string
	//性别
	Sex string
	//最高学历
	HighestEducation string
	//民族
	Nationality string
	//婚姻状况
	MaritalStatus string
	//政治面貌
	PoliticalOutlook string
	//身高
	Height int
	//体重
	Weight int
	//自定义字段
	CustomFields []CustomKVField
}

//自定义字段
type CustomKVField struct {
	//字段名
	FieldName string
	//字段值
	FieldValue string
}

//头像模块
type HeadPortraitCVM struct {
	*TitleCVM
	//头像路径
	PortraitPath string
	//头像风格
	PortraitStyle string
}

//求职意向模块
type JobIntensionCVM struct {
	*TitleCVM
	//意向岗位
	IntentionalPost string
	//职业类型：{全职，兼职，实习，不填}
	CareerType string
	//意向城市
	IntentionalCity string
	//入职时间:{随时，1周内，3个月，另议，不填}
	Hiredate string
	//薪资下限
	SalaryLower float64
	//薪资上限
	SalaryCap float64
	//是否面议
	SalaryPersonally bool
}

//时间段类模块
type TimeRangeCVM struct {
	*TitleCVM
	//时间隐藏
	ExpTimeDisplay bool
	//描述隐藏
	ExpDescDisplay bool
	//经历明细
	ExperienceItems []ExperienceItem
}

//时间经历
type ExperienceItem struct {
	*TitleCVM
	//开始时间
	BeginDate common.JSONTime
	//结束时间
	EndDate common.JSONTime
	//补充字段名一
	ExtFieldName1st string
	//补充字段值一
	ExtFieldVal1st string
	//补充字段名二
	ExtFieldName2nd string
	//补充字段值二
	ExtFieldVal2nd string
	//经验描述
	ExpDesc string
	//显示序号
	viewOrder int64
}

//工作经历
type WorkExperienceCVM struct {
	*TimeRangeCVM
}

//志愿者经历
type VolunteerExperienceCVM struct {
	*WorkExperienceCVM
}

//实习生经历
type TraineeExperienceCVM struct {
	*WorkExperienceCVM
}

//教育经历
type EducateExperienceCVM struct {
	*TimeRangeCVM
}

//项目经历
type ProjectExperienceCVM struct {
	*TimeRangeCVM
}

//自定义经验模块
type CustomExperienceCVM struct {
	*TimeRangeCVM
}

//描述类模块
type DescribeCVM struct {
	*TitleCVM
	//描述
	CVMDescription string
}

//自我评价模块
type SelfEvaluationCVM struct {
	*DescribeCVM
}

//荣誉奖项模块
type HonorsCVM struct {
	*DescribeCVM
}

//自定义描述模块
type CustomDescribeCVM struct {
	*DescribeCVM
}

//自荐信模块
type SelfRecomdLetterCVM struct {
	*DescribeCVM
}

//作品展示模块
type WorksShowCVM struct {
	*TitleCVM
	//作品
	OpuieItems map[string]interface{}
}

//作品
type OpuieItem struct {
	//作品描述
	OpuieItemDesc string
	//作品类型:{"图片","在线"}
	OpuieItemType string
	//访问路径
	ItemVisitPath string
}

//图片类作品
type PictureOpuieItem struct {
	*OpuieItem
	//作品标题
	PictureOpuieTitle string
}

//线上类作品
type OnlineOpuie struct {
	*OpuieItem
}

//标签类模块
type LabelCVM struct {
	*TitleCVM
}

//兴趣爱好模块
type InterestCVM struct {
	*LabelCVM
	InterestLabels []CommonLabel
}

//技能特长模块
type SpecialityCVM struct {
	*LabelCVM
	SpecialityLabels []CommonLabel
}

//通用标签
type CommonLabel struct {
	//标签名称
	LabelName string
	//标签类型
	LableType string
}

//简历默认构造函数
func NewCurriculumVitae() *CurriculumVitae {
	jtnow := common.JSONTime(time.Now())
	cv := &CurriculumVitae{
		CVisibiliType: CV_VISIBILI_TYPE_PRIVATE,
		CVCreateTime:  jtnow,
		CVUpdateTime:  jtnow,
		CVModules:     make(map[string]interface{}),
	}
	//添加基本信息模块
	bicvm := NewBasicInfoCVM()
	cv.CVModules[bicvm.CVMType] = bicvm

	//添加头像模块
	hpcvm := NewHeadPortraitCVM()
	cv.CVModules[hpcvm.CVMType] = hpcvm

	//添加求职意向模块
	jicvm := NewJobIntensionCVM()
	cv.CVModules[jicvm.CVMType] = jicvm

	//添加项目经历模块
	pecvm := NewProjectExperienceCVM()
	cv.CVModules[pecvm.CVMType] = pecvm

	//添加教育经历模块
	eecvm := NewEducateExperienceCVM()
	cv.CVModules[eecvm.CVMType] = eecvm

	//添加自我评价模块
	secvm := NewSelfEvaluationCVM()
	cv.CVModules[secvm.CVMType] = secvm

	return cv
}

//基本信息模块默认构造函数
func NewBasicInfoCVM() *BasicInfoCVM {
	cvm := &BasicInfoCVM{
		TitleCVM: &TitleCVM{
			CVModule: &CVModule{
				CVMType:         CVMTYPE_BASICINFO,
				CVMName:         "基本信息",
				ModuleDisplay:   true,
				ModuleRequisite: true,
			},
			TitleName:    "基本信息",
			TitleDisplay: true,
		},
	}
	return cvm
}

//头像模块默认构造函数
func NewHeadPortraitCVM() *HeadPortraitCVM {
	cvm := &HeadPortraitCVM{
		TitleCVM: &TitleCVM{
			CVModule: &CVModule{
				CVMType:         CVMTYPE_HEADPORTRAIT,
				CVMName:         "头像",
				ModuleDisplay:   true,
				ModuleRequisite: false,
			},
			TitleName:    "头像",
			TitleDisplay: false,
		},
	}
	return cvm
}

//求职意向模块默认构造函数
func NewJobIntensionCVM() *JobIntensionCVM {
	cvm := &JobIntensionCVM{
		TitleCVM: &TitleCVM{
			CVModule: &CVModule{
				CVMType:         CVMTYPE_JOBINTENSION,
				CVMName:         "求职意向",
				ModuleDisplay:   true,
				ModuleRequisite: true,
			},
			TitleName:    "求职意向",
			TitleDisplay: true,
		},
	}
	return cvm
}

//项目经历模块默认构造函数
func NewProjectExperienceCVM() *ProjectExperienceCVM {
	cvm := &ProjectExperienceCVM{
		TimeRangeCVM: &TimeRangeCVM{
			TitleCVM: &TitleCVM{
				CVModule: &CVModule{
					CVMType:         CVMTYPE_PROJECTEXPERIENCE,
					CVMName:         "项目经验",
					ModuleDisplay:   true,
					ModuleRequisite: true,
				},
				TitleName:    "项目经验",
				TitleDisplay: true,
			},
			ExpTimeDisplay: true,
			ExpDescDisplay: true,
		},
	}
	return cvm
}

//教育经历模块默认构造函数
func NewEducateExperienceCVM() *EducateExperienceCVM {
	cvm := &EducateExperienceCVM{
		TimeRangeCVM: &TimeRangeCVM{
			TitleCVM: &TitleCVM{
				CVModule: &CVModule{
					CVMType:         CVMTYPE_EDUCATEEXPERIENCE,
					CVMName:         "教育经历",
					ModuleDisplay:   true,
					ModuleRequisite: true,
				},
				TitleName:    "教育经历",
				TitleDisplay: true,
			},
			ExpTimeDisplay: true,
			ExpDescDisplay: true,
		},
	}
	return cvm
}

//自我评价模块默认构造函数
func NewSelfEvaluationCVM() *SelfEvaluationCVM {
	cvm := &SelfEvaluationCVM{
		DescribeCVM: &DescribeCVM{
			TitleCVM: &TitleCVM{
				CVModule: &CVModule{
					CVMType:         CVMTYPE_SELFEVALUATION,
					CVMName:         "自我评价",
					ModuleDisplay:   true,
					ModuleRequisite: true,
				},
				TitleName:    "自我评价",
				TitleDisplay: true,
			},
		},
	}
	return cvm
}

func main(){
	
}