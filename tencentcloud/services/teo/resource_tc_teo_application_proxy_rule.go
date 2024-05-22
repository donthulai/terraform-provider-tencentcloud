// Code generated by iacg; DO NOT EDIT.
package teo

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	teo "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo/v20220901"

	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func ResourceTencentCloudTeoApplicationProxyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudTeoApplicationProxyRuleCreate,
		Read:   resourceTencentCloudTeoApplicationProxyRuleRead,
		Update: resourceTencentCloudTeoApplicationProxyRuleUpdate,
		Delete: resourceTencentCloudTeoApplicationProxyRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Site ID.",
			},

			"proxy_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Proxy ID.",
			},

			"rule_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Rule ID.",
			},

			"proto": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Protocol. Valid values: `TCP`, `UDP`.",
			},

			"port": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Valid values: `80` means port 80; `81-90` means port range 81-90.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"origin_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Origin server type. Valid values: `custom`: Specified origins; `origins`: An origin group.",
			},

			"origin_port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Origin port, supported formats: single port: 80; Port segment: 81-90, 81 to 90 ports.",
			},

			"origin_value": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Origin site information: When `OriginType` is `custom`, it indicates one or more origin sites, such as `['8.8.8.8', '9.9.9.9']` or `OriginValue=['test.com']`; When `OriginType` is `origins`, there is required to be one and only one element, representing the origin site group ID, such as `['origin-537f5b41-162a-11ed-abaa-525400c5da15']`.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Status, the values are: `online`: enabled; `offline`: deactivated; `progress`: being deployed; `stopping`: being deactivated; `fail`: deployment failure/deactivation failure.",
			},

			"forward_client_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Passes the client IP. Default value is `OFF`. When Proto is TCP, valid values: `TOA`: Pass the client IP via TOA; `PPV1`: Pass the client IP via Proxy Protocol V1; `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the client IP. When Proto=UDP, valid values: `PPV2`: Pass the client IP via Proxy Protocol V2; `OFF`: Do not pass the client IP.",
			},

			"session_persist": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Specifies whether to enable session persistence. Default value is false.",
			},
		},
	}
}

func resourceTencentCloudTeoApplicationProxyRuleCreate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_teo_application_proxy_rule.create")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	var (
		zoneId  string
		proxyId string
		ruleId  string
	)
	var (
		request  = teo.NewCreateApplicationProxyRuleRequest()
		response = teo.NewCreateApplicationProxyRuleResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
	}
	if v, ok := d.GetOk("proxy_id"); ok {
		proxyId = v.(string)
	}

	if v, ok := d.GetOk("zone_id"); ok {
		request.ZoneId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("proxy_id"); ok {
		request.ProxyId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("port"); ok {
		portSet := v.(*schema.Set).List()
		for i := range portSet {
			port := portSet[i].(string)
			request.Port = append(request.Port, helper.String(port))
		}
	}

	if v, ok := d.GetOk("proto"); ok {
		request.Proto = helper.String(v.(string))
	}

	if v, ok := d.GetOk("origin_type"); ok {
		request.OriginType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("origin_port"); ok {
		request.OriginPort = helper.String(v.(string))
	}

	if v, ok := d.GetOk("origin_value"); ok {
		originValueSet := v.(*schema.Set).List()
		for i := range originValueSet {
			originValue := originValueSet[i].(string)
			request.OriginValue = append(request.OriginValue, helper.String(originValue))
		}
	}

	if v, ok := d.GetOk("forward_client_ip"); ok {
		request.ForwardClientIp = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("session_persist"); ok {
		request.SessionPersist = helper.Bool(v.(bool))
	}

	err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
		result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseTeoClient().CreateApplicationProxyRuleWithContext(ctx, request)
		if e != nil {
			if err := resourceTencentCloudTeoApplicationProxyRuleCreateRequestOnError0(ctx, request, e); err != nil {
				return err
			}
			return tccommon.RetryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create teo application proxy rule failed, reason:%+v", logId, err)
		return err
	}

	ruleId = *response.Response.RuleId

	if err := resourceTencentCloudTeoApplicationProxyRuleCreatePostHandleResponse0(ctx, response); err != nil {
		return err
	}

	d.SetId(strings.Join([]string{zoneId, proxyId, ruleId}, tccommon.FILED_SP))

	return resourceTencentCloudTeoApplicationProxyRuleRead(d, meta)
}

func resourceTencentCloudTeoApplicationProxyRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_teo_application_proxy_rule.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	service := TeoService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	idSplit := strings.Split(d.Id(), tccommon.FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	proxyId := idSplit[1]
	ruleId := idSplit[2]

	_ = d.Set("zone_id", zoneId)

	_ = d.Set("proxy_id", proxyId)

	respData, err := service.DescribeTeoApplicationProxyRuleById(ctx, ruleId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `teo_application_proxy_rule` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	if respData.RuleId != nil {
		_ = d.Set("rule_id", respData.RuleId)
	}

	if respData.Proto != nil {
		_ = d.Set("proto", respData.Proto)
	}

	if respData.Port != nil {
		_ = d.Set("port", respData.Port)
	}

	if respData.OriginType != nil {
		_ = d.Set("origin_type", respData.OriginType)
	}

	if respData.OriginPort != nil {
		_ = d.Set("origin_port", respData.OriginPort)
	}

	if respData.OriginValue != nil {
		_ = d.Set("origin_value", respData.OriginValue)
	}

	if respData.Status != nil {
		_ = d.Set("status", respData.Status)
	}

	if respData.ForwardClientIp != nil {
		_ = d.Set("forward_client_ip", respData.ForwardClientIp)
	}

	if respData.SessionPersist != nil {
		_ = d.Set("session_persist", respData.SessionPersist)
	}

	_ = zoneId
	_ = proxyId
	return nil
}

func resourceTencentCloudTeoApplicationProxyRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_teo_application_proxy_rule.update")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	immutableArgs := []string{"zone_id", "proxy_id"}
	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}
	idSplit := strings.Split(d.Id(), tccommon.FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	proxyId := idSplit[1]
	ruleId := idSplit[2]

	needChange := false
	mutableArgs := []string{"proto", "port", "origin_type", "origin_port", "origin_value", "forward_client_ip", "session_persist"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		request := teo.NewModifyApplicationProxyRuleRequest()

		request.ZoneId = &zoneId

		request.ProxyId = &proxyId

		request.RuleId = &ruleId

		if v, ok := d.GetOk("proto"); ok {
			request.Proto = helper.String(v.(string))
		}

		if v, ok := d.GetOk("port"); ok {
			portSet := v.(*schema.Set).List()
			for i := range portSet {
				port := portSet[i].(string)
				request.Port = append(request.Port, helper.String(port))
			}
		}

		if v, ok := d.GetOk("origin_type"); ok {
			request.OriginType = helper.String(v.(string))
		}

		if v, ok := d.GetOk("origin_port"); ok {
			request.OriginPort = helper.String(v.(string))
		}

		if v, ok := d.GetOk("origin_value"); ok {
			originValueSet := v.(*schema.Set).List()
			for i := range originValueSet {
				originValue := originValueSet[i].(string)
				request.OriginValue = append(request.OriginValue, helper.String(originValue))
			}
		}

		if v, ok := d.GetOk("forward_client_ip"); ok {
			request.ForwardClientIp = helper.String(v.(string))
		}

		if v, ok := d.GetOkExists("session_persist"); ok {
			request.SessionPersist = helper.Bool(v.(bool))
		}

		err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
			result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseTeoClient().ModifyApplicationProxyRuleWithContext(ctx, request)
			if e != nil {
				return tccommon.RetryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update teo application proxy rule failed, reason:%+v", logId, err)
			return err
		}
	}

	needChange1 := false
	mutableArgs1 := []string{"status"}
	for _, v := range mutableArgs1 {
		if d.HasChange(v) {
			needChange1 = true
			break
		}
	}

	if needChange1 {
		request1 := teo.NewModifyApplicationProxyRuleStatusRequest()

		request1.ZoneId = &zoneId

		request1.ProxyId = &proxyId

		request1.RuleId = &ruleId

		if v, ok := d.GetOk("status"); ok {
			request1.Status = helper.String(v.(string))
		}

		err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
			result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseTeoClient().ModifyApplicationProxyRuleStatusWithContext(ctx, request1)
			if e != nil {
				return tccommon.RetryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request1.GetAction(), request1.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update teo application proxy rule failed, reason:%+v", logId, err)
			return err
		}
	}

	if err := resourceTencentCloudTeoApplicationProxyRuleUpdateOnExit(ctx); err != nil {
		return err
	}

	return resourceTencentCloudTeoApplicationProxyRuleRead(d, meta)
}

func resourceTencentCloudTeoApplicationProxyRuleDelete(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_teo_application_proxy_rule.delete")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)
	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	idSplit := strings.Split(d.Id(), tccommon.FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	proxyId := idSplit[1]
	ruleId := idSplit[2]

	var (
		request  = teo.NewModifyApplicationProxyRuleStatusRequest()
		response = teo.NewModifyApplicationProxyRuleStatusResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
	}
	if v, ok := d.GetOk("proxy_id"); ok {
		proxyId = v.(string)
	}

	request.ZoneId = &zoneId

	request.ProxyId = &proxyId

	request.RuleId = &ruleId

	status := "offline"
	request.Status = &status

	err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
		result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseTeoClient().ModifyApplicationProxyRuleStatusWithContext(ctx, request)
		if e != nil {
			return tccommon.RetryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create teo application proxy rule failed, reason:%+v", logId, err)
		return err
	}

	_ = response
	if err := resourceTencentCloudTeoApplicationProxyRuleDeletePostHandleResponse0(ctx, response); err != nil {
		return err
	}

	var (
		request1  = teo.NewDeleteApplicationProxyRuleRequest()
		response1 = teo.NewDeleteApplicationProxyRuleResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
	}
	if v, ok := d.GetOk("proxy_id"); ok {
		proxyId = v.(string)
	}

	request1.ZoneId = &zoneId

	request1.ProxyId = &proxyId

	request1.RuleId = &ruleId

	err = resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
		result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseTeoClient().DeleteApplicationProxyRuleWithContext(ctx, request1)
		if e != nil {
			return tccommon.RetryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request1.GetAction(), request1.ToJsonString(), result.ToJsonString())
		}
		response1 = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create teo application proxy rule failed, reason:%+v", logId, err)
		return err
	}

	_ = response1
	return nil
}
